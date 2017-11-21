## 1. mymako.py

```python
# ...
import json
import datetime
from django.utils import timezone
from django.http import HttpResponse
from django.core.serializers.json import DjangoJSONEncoder
class JSONEncoder(DjangoJSONEncoder):
    """
    DjangoJSONEncoder subclass that knows how to encode date/time in custom format
    """

    def default(self, o):
        # See "Date Time String Format" in the ECMA-262 specification.
        if isinstance(o, datetime.datetime):
            if timezone.is_aware(o):
                # translate to time in local timezone
                o = timezone.localtime(o)
            return o.strftime('%Y-%m-%d %H:%M:%S %z')
        else:
            return super(DjangoJSONEncoder, self).default(o)


def render_json(dictionary={}):
    '''
    return the json string for response
    @summary: dictionary也可以是string, list数据
    @note:  返回结果是个dict, 请注意默认数据格式:
                                    {'result': '',
                                     'message':''
                                    }
    '''

    if type(dictionary) is not dict:
        # 如果参数不是dict,则组合成dict
        dictionary = {
            'result': True,
            'message': dictionary
        }
    return HttpResponse(json.dumps(dictionary, cls=JSONEncoder), content_type='application/json')
# ...
```

## 2. default.py

```python
# ==============================================================================
# Middleware and apps
# ==============================================================================
MIDDLEWARE_CLASSES = (
    ...
    'common.middlewares.TimezoneMiddleware',  # 时区切换中间件
    'django.middleware.locale.LocaleMiddleware',
)
# ==============================================================================
# timezone
# ==============================================================================
# 开启时区支持
USE_TZ = True
# 默认时区为东八区 +0800
TIME_ZONE = 'Asia/Shanghai'

# ==============================================================================
# i18n
# ==============================================================================
USE_I18N = True
USE_L10N = True

# session失效 ->accept-language -> LANGUAGE_CODE
# 故应设置LANGUAGE_CODE为基础语言(中文)，效果就是基础语言(中文)无需翻译，也就是zh_CN下的po文件可以保留msgstr为空
# 若基础语言为中文，LANGUAGE_CODE设为英文，当某些中文翻译为空时，会导致django去LANGUAGE_CODE指定的文件中寻找翻译内容
# 如果LANGUAGE_CODE指定为英文，则会出现中英文同时存在的情况
LANGUAGE_CODE = 'zh-CN'
# 设定使用根目录的locale
LOCALE_PATHS = (os.path.join(PROJECT_ROOT, 'locale'),)

# 界面可选语言
_ = lambda s: s
LANGUAGES = (
    ('en', _(u'English')),
    ('zh-cn', _(u'简体中文')),
)

LANGUAGE_SESSION_KEY = 'blueking_language'
LANGUAGE_COOKIE_NAME = 'blueking_language'
```

## 3. common.middlewares

```python
# 通过session传递时区
class TimezoneMiddleware(object):
    def process_request(self, request):
        tzname = request.session.get('blueking_timezone')
        if tzname:
            timezone.activate(pytz.timezone(tzname))
        else:
            timezone.deactivate()
```
## 4. 业务切换
```python
# 这个逻辑要写到你的业务切换逻辑中
def change_business(request):
    """
    业务切换
    """
    ...
    # 切换业务的同时切换时区
    timezone = request.POST.get('timezone')
    request.session['blueking_timezone'] = timezone
    ...
```
说明：app本身需要根据自身的设计来调整，cc接口中get_app_by_user、get_app_list等接口均在业务层级增加了
时区和语言字段，目前接口已经部署到dev环境：
```python
{
    ... 
    u'TimeZone': u'Asia/Shanghai', 
    u'Language': u'zh-cn', 
    u'ApplicationID': u'1', 
    ...
}
```
可以有很多种方式来设计：
 - 比如前端加载业务列表的同时，暂存每个业务的时区信息，切换业务的同时传递到后台
 - 又比如后台缓存业务对应的时区信息，切换业务的时候，后台自动匹配时区

## 5. 接口、业务逻辑相关调整
```python
# 工具函数
from django.utils import timezone

def strftime_local(aware_time, fmt="%Y-%m-%d %H:%M:%S %z"):
    """格式化aware_time为本地时间"""

    if timezone.is_aware(aware_time):
        return timezone.localtime(aware_time).strftime(fmt)

    return aware_time.strftime(fmt)

```
1. json返回时间

```python
    {
        'create_time': sometime.strftime('%Y-%m-%d %H:%M:%S') -> sometime
    }
```
    并且使用改造过的`render_json`返回（DRF方式可参考：http://km.oa.com/group/18567/articles/show/323549）
    
2. 模板中使用到时间
    格式化方式修改：
    ```python
        sometime.strftime('%Y-%m-%d %H:%M:%S') -> strftime_local(sometime)
    ```
    
3. 给model中的时间字段赋值

    ```python
        obj = SomeModel.objects.get(pk=1)
        obj.create_time = datetime.datetime.now() + datetime.timedelta(days=1)
   ```
    `改成`
    
    ```python
    from django.utils import timezone    
    obj.create_time = timezone.now() + datetime.timedelta(days=1)
    ```
4. 从第三方系统或前端传入后台的时间
    
    根据业务逻辑酌情处理

## 6. 后台任务调整

后台任务需要调用接口获取业务相关的时区信息，并根据业务逻辑调整从第三方系统获取的时间

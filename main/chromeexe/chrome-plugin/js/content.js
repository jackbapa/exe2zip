var nativeKey = 'nativeMessage';

chrome.runtime.onMessage.addListener(function (request) {
//alert('content.js:addListener--begin');
    //console.log('Content.js ' + JSON.stringify(request));
    var type = request.type;
    var data = request.data;

    //alert('content.js:addListener--type='+type);
    //alert("content.js:addListener--data="+JSON.stringify(data));

    if (type === nativeKey && data) {
        //alert('content.js,inside postMessage');
        window.postMessage({
            data: data,
            type: 'CAContentScript',
        }, '*');
    }
});

// 向页面注入JS
function injectCustomJs(jsPath) {
    jsPath = jsPath || 'js/inject.js';
    var temp = document.createElement('script');
    temp.setAttribute('type', 'text/javascript');
    temp.src = chrome.extension.getURL(jsPath);
    temp.onload = function () {
        this.parentNode.removeChild(this);
    };
    document.documentElement.appendChild(temp);
}

function onInjectMessage() {
    window.addEventListener('message', (e) => {
        var type = e.data.type;
        var action = e.data.action;
        if (type === 'CAInject' && action === 'logout') {
            console.log('CA Certification[CS]: Trigger close');
            chrome.runtime.sendMessage({
                type: 'closeTab',
            });
        }
    }, false);
}


document.addEventListener('NativeCallEvent', function (evt) {
    //alert('content event start='+evt.detail.api);
    console.log('content event start');
    if (typeof chrome.app.isInstalled !== 'undefined') {
        chrome.runtime.sendMessage({
            type: nativeKey,
            api: evt.detail.api,
            domain: evt.detail.domain
        }, function (response) {
            return true;
        });
    }
}, false);

// hijack logout
function hijackLogout(id) {
    document.addEventListener('click', function (e) {
        if (e.target.id && e.target.id === id) {
            console.log('CA Certification[CS]: Trigger close');
            chrome.runtime.sendMessage({
                type: 'closeTab',
            });
        }
    });
}


function sendNativeMessageToBg() {
    var origin = document.location.origin;
    // except: api = origin + base + '/ca/login';
    var api = CaDom.getAttribute('base') || '';
    // base contains ca/login
    if (api.indexOf('ca/login') === -1) {
        api = api + '/ca/login';
        api = api.replace('//ca/login', '/ca/login');
    }
    // base is absolute path
    if (api.indexOf('http') === -1) {
        api = origin + api;
    }
    console.log('CA Certification[CS]: Send bg message ' + JSON.stringify({
        domain: origin,
        api,
    }));
    chrome.runtime.sendMessage({
        type: nativeKey,
        domain: origin,
        api,
    });
}


var CaDom = document.getElementById('__CA__');

if (CaDom) {
    //alert('begin to inject');
    injectCustomJs();
    sendNativeMessageToBg();
    hijackLogout(CaDom.getAttribute('logout-id') || 'logout');
    onInjectMessage();
}

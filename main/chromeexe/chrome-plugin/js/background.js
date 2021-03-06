var nativeKey = 'nativeMessage';
var port;
var mocking = false;
var UserName = 'test';
var UserId = '28db2e05f7e54c3b90a9e41e5cfe2318';
var Authorization;
var Token;

var port = null;
var nativeHostName = "com.exe.name"

// Init Native Message
function connect() {
    port = chrome.runtime.connectNative(nativeHostName);
    port.onMessage.addListener(onNativeMessage);
}

function onNativeMessage(message) {
    //alert('received message from native exe='+JSON.stringify(message));
    //alert('received message from native exe='+message);
    try {
        var messageObj = message || {};
        var success = messageObj.success;
        var result = messageObj.result;
        if (String(success) === 'true') {
            chrome.browserAction.setIcon({path: 'img/success.png'});
            //alert('received message from native exe='+true);
            sendMessageToCS(result);
        } else {
            chrome.browserAction.setIcon({path: 'img/error.png'});
        }
    } catch (error) {
        chrome.browserAction.setIcon({path: 'img/error.png'});
    }
}

chrome.runtime.onMessage.addListener(function (request, sender) {
    var requestObj = request || {};
    var type = requestObj.type;
    var api = requestObj.api;
    var domain = requestObj.domain;

    //alert('background1-type='+type);
    //alert('background1-api='+api);
    //alert('background1-domain='+domain);

    if (['default', 'success', 'error'].indexOf(type) >= 0) {
        chrome.browserAction.setIcon({path: 'img/' + type + '.png'});
    } else if (type === nativeKey) {
        if (mocking) {
            setTimeout(function () {
                chrome.browserAction.setIcon({path: 'img/success.png'});
                sendMessageToCS({
                    Data: {
                        UserName: UserName,
                        UserId: UserId,
                        Authorization: Authorization,
                    },
                    data: {
                        userName: UserName,
                        userId: UserId,
                        token: Authorization,
                    },
                    Token: Token,
                });
            }, 300);
        } else {
            connect();
            sendNativeMessage({
                domain: domain,
                api: api,
            });
        }
    } else if (type === 'closeTab') {
        chrome.tabs.remove(sender.tab.id);
    }
});


function sendNativeMessage(message) {
    if (!port) return;
    port.postMessage(message);
}

function sendMessageToCS(message) {
    chrome.tabs.query({active: true, currentWindow: true}, function (tabs) {
        chrome.tabs.sendMessage(tabs[0].id, {
            type: nativeKey,
            data: message,
        });
    });
}

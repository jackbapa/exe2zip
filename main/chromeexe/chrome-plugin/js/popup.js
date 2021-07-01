var nativeHostName = "com.exe.name"
var port = null;

function appendMessage(text) {
    document.getElementById('caResponse').innerHTML += "<p>" + text + "</p>";
}

function sendNativeMessage() {
    if (!port) return;
    var domain = document.getElementById('caDomain').value;
    var api = document.getElementById('caApi').value;
    var message = {
        domain: domain,
        api: api,
    };
    port.postMessage(message);
    appendMessage("Sent message: " + JSON.stringify(message));
}

function onNativeMessage(message) {
    try {
        var messageObj = message || {};
        var success = messageObj.success;
        var result = messageObj.result;
        appendMessage("Received message: " + 'success is' + success);
        appendMessage("Received message: " + 'result is' + JSON.stringify(result));
    } catch (error) {
        console.error('CA Certification[PP]: ', error);
    }
}

function onDisconnected() {
    appendMessage("Failed to connect: " + chrome.runtime.lastError.message);
}

// Native Message
function connect() {
    port = chrome.runtime.connectNative(nativeHostName);
    port.onMessage.addListener(onNativeMessage);
    port.onDisconnect.addListener(onDisconnected);
}

window.onload = function () {
    // CA TEST
    var caTest = document.getElementById("caTest");
    if (caTest) {
        caTest.addEventListener("click", function (e) {
            connect();
            sendNativeMessage();
        });
    }


    // Mock
    var bg = chrome.extension.getBackgroundPage();
    var ipSwitch = document.getElementById("switch");
    var userWrapper = document.getElementById("user");
    var testWrapper = document.getElementById("test");
    var setInsArr = ["UserName", "UserId", "Authorization", "Token"];

    function setInsValue() {
        setInsArr.forEach(function (ele) {
            if (bg[ele]) {
                var eleInput = document.getElementById(ele);
                eleInput.value = bg[ele];
            }
        });
    }

    if (bg.mocking) {
        ipSwitch.click();
        userWrapper.style.display = "block";
        testWrapper.style.display = "none";
        setInsValue();
    } else {
        userWrapper.style.display = "none";
        testWrapper.style.display = "block";
    }

    ipSwitch.addEventListener("click", function (e) {
        if (e.target.checked) {
            bg.mocking = true;
            userWrapper.style.display = "block";
            testWrapper.style.display = "none";
            setInsValue();
        } else {
            bg.mocking = false;
            userWrapper.style.display = "none";
            testWrapper.style.display = "block";
        }
    });

    var setBtn = document.getElementById("set");
    setBtn.addEventListener("click", function (e) {
        setInsArr.forEach(function (ele) {
            var eleInput = document.getElementById(ele);
            bg[ele] = eleInput.value;
        });
    });
}

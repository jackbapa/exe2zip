var cookie = {
  set(key, value, exdays = 1) {
    if (key && value) {
      var d = new Date();
      d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
      var expires = d.toGMTString();
      document.cookie = '' + key + '=' + encodeURIComponent(value) + ';' + 'expires=' + expires + ';' + 'path=/';
    }
  },
};

function isObject(obj) {
  return Object.prototype.toString.call(obj) === '[object Object]';
}

console.log('CA Certification[IJ]: Successfully injected script');

window.addEventListener('message', (e) => {
  var type = e.data.type;
  var data = e.data.data;
  
  if (type === 'CAContentScript') {
    console.log('CA Certification[IJ]: Received cs message', data);
    var dataObj = data;
    if (typeof data === 'string') {
      try {
        dataObj = JSON.parse(data);
		console.log('inject.js,dataObj=', dataObj);
      } catch (error) {
        console.error('CA Certification[IJ]: Parse result error', error);
      }
    }

    if (dataObj) {
      var Token = dataObj.Token;
      var Data = dataObj.Data;
      var result = dataObj.result;
      var innerData = dataObj.data;
      var ca_error = dataObj.ca_error;

	  console.log('inject.js,dataObj.Token=', dataObj.Token);
	  console.log('inject.js,dataObj.Data=', dataObj.Data);
	  console.log('inject.js,dataObj.result=', dataObj.result);
	  console.log('inject.js,dataObj.data=', dataObj.data);
	  console.log('inject.js,dataObj.ca_error=', dataObj.ca_error);

      if (!Token && !Data && !result && !innerData) {
        console.error('CA Certification[IJ]: No valid user info or new access project？');
        return;
      }

      var c_map = {
        UserName: 'userName',
        UserId: 'userId',
        Authorization: 'authorization',
      };
      var c_keys = Object.keys(c_map);
	  console.log('injected,c_keys='+c_keys);

      if (ca_error) {
        cookie.set('x_ca_error', encodeURI(ca_error));
        console.log('injected, ca_error is not null');
      }

      // QZ
      if (Token) {
        cookie.set('qz_t', Token);
		console.log('injected, Token is not null');
      }
      // NUIMS
      if (Data && isObject(Data)) {
        Object.keys(Data).forEach((key) => {
          var value = Data[key];
		  console.log('injected, Data: value='+value);
          if (c_keys.indexOf(key) >= 0 && value) {
            cookie.set(c_map[key], value);
          }
        });
      }
      // LCP
      if (result && isObject(result)) {
        Object.keys(result).forEach((key) => {
          var value = result[key];
          if (c_keys.indexOf(key) >= 0 && value) {
            cookie.set(c_map[key], value);
          }
        });
      }
      // MH
      if (innerData && isObject(innerData)) {
        var mc_map = {
          token: 'm_t',
          userId: 'm_ud',
          userName: 'm_un',
        };
        var mc_keys = Object.keys(mc_map);
        Object.keys(innerData).forEach((key) => {
          var value = innerData[key];
          if (mc_keys.indexOf(key) >= 0 && value) {
            cookie.set(mc_map[key], value);
          }
        });
      }

      var CaDom = document.getElementById('__CA__');
      if (CaDom) {
        CaDom.setAttribute('init', 'done');
        var initCallback = window.__CA_CB__;
        if (initCallback && typeof initCallback === 'function') {
          initCallback();
        }
      }
    } else {
      console.error('CA Certification[IJ]: No valid cs message');
    }
  }
}, false);

if(!window.__CA_LOGOUT__) {
  window.__CA_LOGOUT__ = function() {
    window.postMessage({
      type: 'CAInject',
      action: 'logout',
    }, '*');
  };
}

var MEDICAL_HEALTH_CONF = {};
var SXKUI_DOMAIN = '39.106.127.158:8081';
var SXKUI_TOKEN  = 'test_token';

// 新首页接口
var sxk_index = {};
sxk_index.API = '//127.0.0.1:8081/';
sxk_index.URL = '//127.0.0.1:8081/templates/';

// 所有接口
var medical_health_all = {};
medical_health_all.API = '//trade.suixingkang.com/test/api/';
medical_health_all.URL = '//trade.suixingkang.com/test/h5/';// Import conf/conf.js // 旧接口


// 新首页
sxk_index.home    = sxk_index.URL+'index.html#/tab/home';
sxk_index.service = sxk_index.URL+'index.html#/tab/service';
sxk_index.my      = sxk_index.URL+'index.html#/tab/my';
sxk_index.show365 = sxk_index.URL+'asset/card365.html';
sxk_index.searchcard  = sxk_index.URL+'index.html#/searchCard';
MEDICAL_HEALTH_CONF.index = sxk_index


var MEDICAL_HEALTH_AJAX = function(obj) {
	
	// 异常处理
	var _theEcp = new MEDICAL_HEALTH_UI_ECP(obj);
	
	// JSON处理
	this.json = function(url, param, callback, request) {
		$.ajax({
	        type: request,
	        url: url,
	        async: false,
	        cache: false,
	        dataType: "text",
	        data: param,
	        headers: {'Sxk-Agent':'sxkui1.1'},
	        dataType: "json",
			success: function(data) {
				if (_theEcp.check(data, url)){
					callback(data);
				} 
			}
		});
	}
	
	// JSONP处理,访问其他接口
	this.jsonp = function(url, param, callback, request) {
		$.ajax({
	        type: request,
	        url: url,
	        async: false,
	        cache: false,
	        dataType: "text",
	        data: param,
	        headers: {'Sxk-Agent':'sxkui1.1'},
	        dataType: "jsonp",
			jsonp:"callback",
			success: function(data) {
				if (_theEcp.check(data, url)){
					callback(data);
				}
			}
		});
	}
	
	// 文件上传 依赖ajaxUPload
	this.upload = function(url, param, callback) {
		$.ajaxFileUpload({
			url: url,
			type: 'post',
			secureuri: false,
			fileElementId: param.fileId,
			dataType: 'json',
			data: param,
			headers: {'Sxk-Agent':'sxkui1.1'},
			success: function(data, status) {
				if (_theEcp.check(data, url)){
					callback(data);
				}
			},
			error: function(data, status, e) {
				//alert(e);
			}
	   });
	}
}

var MEDICAL_HEALTH_UI_ENV = function() {
	// 是不是ANDROID
	this.isAndroid = function (){
	    var ua = window.navigator.userAgent.toLowerCase();
	    if(ua.match(/MEDICAL_HEALTH_ANDROID/i) == 'medical_health_android'){
	        return true;
	    }else{
	        return false;
	    }
	}
	
	// 是不是IOS
	this.isIOS = function (){
	    var ua = window.navigator.userAgent.toLowerCase();
	    if(ua.match(/MEDICAL_HEALTH_IOS/i) == 'medical_health_ios'){
	        return true;
	    }else{
	        return false;
	    }
	}
	
    // 是不是微信
	this.isWeixin = function (){
	    var ua = window.navigator.userAgent.toLowerCase();
	    if(ua.match(/MicroMessenger/i) == 'micromessenger'){
	        return true;
	    }else{
	        return false;
	    }
	}
}
// 全局Ajax请，依赖jQuery
var MEDICAL_HEALTH_UI_AJAX = function(obj) {
	
	// 异常处理
	var _theEcp = new MEDICAL_HEALTH_UI_ECP(obj);
	
	// JSON处理
	this.json = function(url, param, callback, request) {
		$.ajax({
	        type: request,
	        url: url,
	        async: false,
	        cache: false,
	        dataType: "text",
	        data: param,
	        headers: {'Sxk-Agent':'sxkui1.1'},
	        dataType: "json",
			success: function(data) {
				if (_theEcp.check(data, url)){
					callback(data);
				} 
			}
		});
	}
	
	// JSONP处理,访问其他接口
	this.jsonp = function(url, param, callback, request) {
		$.ajax({
	        type: request,
	        url: url,
	        async: false,
	        cache: false,
	        dataType: "text",
	        data: param,
	        headers: {'Sxk-Agent':'sxkui1.1'},
	        dataType: "jsonp",
			jsonp:"callback",
			success: function(data) {
				if (_theEcp.check(data, url)){
					callback(data);
				}
			}
		});
	}
	
	// 文件上传 依赖ajaxUPload
	this.upload = function(url, param, callback) {
		$.ajaxFileUpload({
			url: url,
			type: 'post',
			secureuri: false,
			fileElementId: param.fileId,
			dataType: 'json',
			data: param,
			headers: {'Sxk-Agent':'sxkui1.1'},
			success: function(data, status) {
				if (_theEcp.check(data, url)){
					callback(data);
				}
			},
			error: function(data, status, e) {
				//alert(e);
			}
	   });
	}
}
// Import comm/exception.js // 错误处理
var MEDICAL_HEALTH_UI_ECP = function(obj) {
	
	// 处理类
	this.OBJ = obj;
	
	// 白名单:不需
	this.White = obj.whiteList;
	
	// 数据判断
	this.check = function(data, url) {
		// 白名单
		for(var i in this.White) {
			var action = this.White[i];
			if (url.indexOf(action) != -1) {
				return true;
			}
		}
		
		// 未登录TOKEN
		if (data.login == false) {
			this.login();
			return false;
		}
		
		// 未激活
		if (data.access == false) {
			this.active();
			return false;
		}
		
		return true;
	}
	
	// 登录
	this.login = function(data) {
		if (callback = this.OBJ.exback['login']) {
			callback(data);
		} else {
			this.OBJ.login();
		}
	}
	
	// 未激活
	this.active = function(data) {
		if (callback = this.OBJ.exback['active']) {
			callback(data);
		} else {
			this.OBJ.buy365();
		}
	}
}
// Import env/wap.js 
// WAP 通用情况使用
var SXKUI_WAP  = function(config) {
	
	// 全局配置
	this.config = config;
	this.exback = {};
	this.whiteList = {};
	
	// 配置自定义异常
	this.exception = function(name, callback) {
		this.exback[name] = callback;
	}
	
	// 配置白名单
	this.white = function(white) {
		if (white != '') {
			this.whiteList = white.split(',');
		}
	}
	
	// 当前环境
	this.env = this.config.env;
		
	// SSO & Cookie
	this.setCookie = function(name, value, Days) {
		var exp = new Date();
		exp.setTime(exp.getTime() + Days*24*60*60*30);
		document.cookie = name + "="+  (value) + ";expires=" + exp.toGMTString() + ";path=/;domain=" + SXKUI_DOMAIN;
		//console.log(value);
	}
	this.getCookie = function(name) {
			var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
			if(arr=document.cookie.match(reg))
			return unescape(arr[2]);
		else {
		 	return null;
		}
	}	
		
	// TOKEN
	this.setToken = function(token) {
		this.setCookie(SXKUI_TOKEN, token);
	}
	this.getToken = function() {
		
		// 如果配置中有Token
		if (token = this.config.token) return token;
		
		// 新系统取法
		if(token = this.getQuery('token')){
			this.setToken(token);
		}else{
			token = this.getCookie(SXKUI_TOKEN);
		}
		if (token) return token;
		
		// 兼容storage方式
		if (token = localStorage.getItem('token')) {
			return token;
		}
		
		// 兼容旧系统获取方法
		var data = eval('('+localStorage.getItem('data')+')');
		try{
			this.setToken(data.uid);
			return data.uid 
		}catch(e){
			return '';
		}
		return '';
	}
		
	// 缓存处理全局
	// 设置缓存
	this.getCache = function(key) {
		return localStorage.getItem('_cache_'+key);
	}
	// 获取缓存
	this.setCache = function(key, value) {
		return localStorage.setItem('_cache_'+key, value);
	}
	// 获取参数
	this.getParam = function(key) {
		if (value = this.getQuery(key)) {
			this.setCache(key, value);
			return value;
		} else {
			return this.getCache(key);
		}
	}
		
	// 设置标题
	this.setTitle = function(title)	{
		document.title = title;
	}
	
	// 关闭webview
	this.closeWebView = function() {
		
	}
	// webview到首页
	this.homeWebView = function() {
		
	}
	
	// 获取API
	this.getAPI = function() {
		var action = arguments[0] || '';
		return this.config.API + action;
	}
	
	// 获取URL
	this.getURL = function() {
		var action = arguments[0] || '';
		return this.config.URL + action;
	}
	
	// 跳转指定URL
	this.goURL = function(url) {
		location.href = this.getURL(url);
	}
	
	// 返回首页
	this.home = function(){
		this.activity('index.home');
	}
	
	// 登录页
	this.login = function(){
		this.activity('user.login');
	}
	
	// 激活
	this.active = function() {
		this.activity('user.active');
	}
	
	// 购买365用户
	this.buy365 = function() {
		this.activity('www.buy365');
	}
	
	// 模块分析
	this.getModule = function(index) {
		var indexs = index.split('.');
		if (indexs.length == 1) {
			var module = this.config.module;
		} else {
			var module = indexs[0];
			var index = indexs[1];
		}
		var url = MEDICAL_HEALTH_CONF[module][index] || '';
		return {module:module, url:url, index:index};
	}
	
	// 获取Query值
	this.getQuery = function (name) { 
	    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i"); 
	    var r = window.location.search.substr(1).match(reg); 
	    if (r != null) return unescape(r[2]); 
	    return null; 
	}

	// 统计代码
	this.tongji = function() {
		(function() {
		  var hm = document.createElement("script");
		  hm.src = "https://hm.baidu.com/hm.js?0fbfbad0747eb55c93cce5d06f116504";
		  var s = document.getElementsByTagName("script")[0]; 
		  s.parentNode.insertBefore(hm, s);
		})();
	}

	// 百度统计
	// https://www.jianshu.com/p/40cc4a1f3806
	this.baidu = function(path) {
		var _hmt = window._hmt || [];
		_hmt.push(['_setAutoPageview', false]);
		_hmt.push(['_trackPageview', '/h5/index'+path]);
	}

	// 安全模式
	var safe = this.config.safe || 'ON';
	if (safe == 'ON') {
		var token = this.getToken();
		var url = window.location.href;
		if (url.indexOf('token') != -1) {
			url = url.replace('token='+token, '');
			history.replaceState({}, '', url);
		}
		this.tongji();
	}
	
	// 隐藏底部导航
	this.tabbar = function() {
		//
	}
	
	// 返回指定页
	this.activity = function(index) {
		var data = this.getModule(index);
		// 如果是WWW跳转
		if (data.module == 'www' && MEDICAL_HEALTH_CONF.www.noproxy.indexOf(data.index) == -1) {
			//data.url = MEDICAL_HEALTH_CONF.www.proxy.replace('{uri}', encodeURIComponent(data.url));
			data.url = MEDICAL_HEALTH_CONF.www.proxy.replace('{uri}', data.url);
		}
		// 兼容测试模式
		if (document.domain == 'www.ji.com') {
			var arr = data.url.split('#');
			var url = arr[0];
			var uri = arr[1];
			url = url.indexOf('?') == -1 ? url + '?token={token}' : url + '&token={token}';
			uri = uri ? '#' + uri : '';
			data.url = url + uri;
		}
		var params = arguments[1] || {};
		// 如果接口跳转
		if (!params._apiUrl) {
			params.token = this.getToken();
			for(var key in params) {
				data.url = data.url.replace("{"+key+"}", params[key]);
			}
			data.url = data.url.replace(/\{.*?\}/g, '');
			location.href = data.url;
		} else {
			var _apiUrl = params._apiUrl;
			this.json(index,function(data){
				location.href = data.datas[_apiUrl];
			});
		}
	}
	
	// JSON提交
	this.json = function() {
		if (arguments.length == 2) {
			var action   = arguments[0];
			var param    = {};
			var callback = arguments[1]
		} else {
			var action   = arguments[0];
			var param    = arguments[1] || {};
			var callback = arguments[2]
		}
		// Action分析
		param.token = this.getToken();
		var request = param.method || 'POST';
		delete param.method;
		var Ajax = new MEDICAL_HEALTH_AJAX(this);
		var actions = action.split('.');
		if (actions.length == 1) {
			var url = this.getAPI(action);
			Ajax.json(url, param, callback, request);
		} else {
			var url = MEDICAL_HEALTH_CONF[actions[0]]['API'] + actions[1];
			Ajax.jsonp(url, param, callback, request);
		}
	}
	
	// 文件上传
	this.upload = function(action, param, callback) {
		var request = param.method || 'POST';
		var url = this.getAPI(action);
		param.token = this.getToken();
		delete param.method;
		var Ajax = new MEDICAL_HEALTH_AJAX(this);
		Ajax.upload(url, param, callback, request);
	}
	
	// base64上传
	this.base64 = function(action, param, callback) {
		var __this = this;
		var fileId = param.fileId;
		var body   = param.base64;
		var file = $('#'+fileId).get(0).files[0];
		var reader = new FileReader(); 
		reader.readAsDataURL(file); 
		reader.onload = function(e){ 
 			param[body] = this.result;
 			delete param.fileId;
		    delete param.base64;
		    console.log(param);
		    __this.json(action, param, callback);
	   }
	}
	
	// 插件扩展
	this.plugin = function(name) {
		var OBJ = eval('new SXKUI_PLUGIN_'+name);
		OBJ.setConf(this);
		return OBJ;
	}
}// Import env/weixin.js // 微信  环境


var SXKUI_Weixin = function(config) {
	
	// 初始化基类
	var _this = new SXKUI_WAP(config);
	
	return _this;
}// Import env/ios.js // IOS 环境
var SXKUI_IOS = function(config) {
	
	// 初始化基类
	var _this =  new SXKUI_WAP(config);
	var _parent = {};
	$.extend(_parent, _this);
	
	// 设置首页
	_this.setTitle = function(title) {
		try {	
			window.onpageshow = function(event) {　　
				if(event.persisted) {　　　　
					window.webkit.messageHandlers.setTitle.postMessage({title:title});　
				}
			};
			window.webkit.messageHandlers.setTitle.postMessage({title:title});
		}catch(e){
           //
		}
	}
	
	// 跳转首页
	_this.home = function() {
		window.webkit.messageHandlers.goHome.postMessage(null);
	}
	
	// 关闭webview
	_this.closeWebView = function() {
		window.webkit.messageHandlers.closeCurPage.postMessage(null);
	}
	// webview到首页
	_this.homeWebView = function() {
		
	}
	
	// 登录页
	_this.login = function(){
		window.webkit.messageHandlers.goLogin.postMessage(null);
	}
	
	// 激活页
	_this.active = function() {
		window.webkit.messageHandlers.goBind365.postMessage(null);
	}
	
	// 关闭底部导航 
	_this.tabbar = function(flag) {
		var hidden = !flag;
		window.webkit.messageHandlers.setTabbarHidden.postMessage({'hidden' : hidden});
		if (callback = arguments[1]){
			callback();
		} 		
	}
	
	// 跳转
	_this.activity = function(index) {
		var config = _this.getModule(index);
		var module = config.module;
		var action = config.index.replace('/init', '');
		var params = arguments[1] || {};
		params.token = _this.getToken();
		
		if (MEDICAL_HEALTH_CONF.IOS[module] && MEDICAL_HEALTH_CONF.IOS[module][action]) {
			MEDICAL_HEALTH_CONF.IOS[module][action](index, params, _this);
		}else {
			_parent.activity(index, params);
		}
	}
	
	return _this;
}// Import env/android.js // Android 环境
var SXKUI_Android = function(config) {
	
	// 初始化基类
	var _this = new SXKUI_WAP(config);
	var _parent = {};
	$.extend(_parent, _this);
	
	// 设置首页
	_this.setTitle = function(title) {
		try{
			window.android.showAppTitle(title);
		}catch(e){}
	}
	
	// 关闭webview
	_this.closeWebView = function() {
		try{
			window.android.closeCurActivity();
		}catch(e){}
	}
	// webview到首页
	_this.homeWebView = function() {
		window.location.href="protocol://android?params="+JSON.stringify({"action": 2});
	}
	
	// 跳转首页
	_this.home = function() {
		//console.log('android: home');
		window.location.href="protocol://android?params="+JSON.stringify({"action": 1});
	}
	
	// 登录页
	_this.login = function(){
		//console.log('android: login');
	}
	
	// 激活页
	_this.active = function() {
		//console.log('android: active');
	}
	
	// 跳转
	_this.activity = function(index) {
		var config = _this.getModule(index);
		var module = config.module;
		var action = config.index.replace('/init', '');
		var params = arguments[1] || {};
		params.token = _this.getToken();
		
		if (MEDICAL_HEALTH_CONF.Android[module] && MEDICAL_HEALTH_CONF.Android[module][action]) {
			MEDICAL_HEALTH_CONF.Android[module][action](index, params, _this);
		}else {
			_parent.activity(index, params);
		}
	}
	
	return _this;
}
// Import plugin/floatHome.js // 插件:浮动返回首页
var SXKUI_PLUGIN_floatHome = function(){
	
	// 环境实例
	var _ENV = {};
	
	// 环境方法
	this.setConf = function(ENV) {
		_ENV  = ENV;
	}
	
	// 显示
	this.show = function() {
		$('#SXKUI_floatHome_button').show();
		return this;
	}
	
	// 隐藏
	this.hide = function() {
		$('#SXKUI_floatHome_button').hide();
		return this;
	}
	
	// CSS方法
	this.css = function() {
		return "position: fixed;" + 
	    "right: 38px;" +
	    "bottom: 20px;" +
	    "z-index: 1000000002;" +
	    "width: 40px;" +
	    "height: 40px;" +
	    "border-radius: 50%;" +
	    "text-align: center;" +
	    "background: url(data:image/png;base64,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) no-repeat;" +
	    "background-size: 100% 100%;" +
	    "display: flex;" +
	    "justify-content: space-around;"
	}
	
	// 初始化
	this.init = function() {
		$("<div id='SXKUI_floatHome_button' style='"+this.css()+"'></div>").on('click', function(){
			_ENV.home();
		}).hide().appendTo($("body"));
	};
	this.init();
};// Import plugin/reback.js // 插件：循环跳转的问题
var SXKUI_PLUGIN_reback = function(){
		
	// 环境实例
	var __this = this;
	var _ENV = {};
	
	// 环境方法
	this.setConf = function(ENV) {
		_ENV  = ENV;
	}
		
	// 初始化
	this.init = function() {
		var isBridge = localStorage.getItem('isBridge');
		if(!isBridge) {
			localStorage.setItem('isBridge', 'ok');
		} else {
			localStorage.removeItem('isBridge');
			history.back();
		}
	};
	
	// 兼容IOS
	window.onpageshow = function(event) {　　
		if(event.persisted) {　　　　
			__this.init();　　
		}
	};
	
	this.init();
};// Import plugin/weixin.js // 插件：微信分享
var SXKUI_PLUGIN_weixin = function(){
		
	// 环境实例
	var __this = this;
	var __ENV = {};
	var shareData = {};
	
	// 环境方法
	this.setConf = function(ENV) {
		__ENV  = ENV;
	}
				
	// 自定义“分享到朋友圈”及“分享到QQ空间”按钮的分享内容
	this.wxRead = function(data) {
	
		wx.ready(function () {
			// onMenuShareTimeline
			wx.onMenuShareTimeline({
		        title: data.title,
		        link: data.link,
		        imgUrl: data.imgUrl,
			    success: function () {}
			});
			
			// onMenuShareAppMessage
			wx.onMenuShareAppMessage({
		        title: data.title,
		        desc: data.desc,
		        link: data.link,
		        imgUrl: data.imgUrl,
				type: 'link',
				dataUrl: '',
				success: function () {}
			});
			
			// onMenuShareQQ
			wx.onMenuShareQQ({
		        title: data.title,
		        desc: data.desc,
		        link: data.link,
		        imgUrl: data.imgUrl,
				success: function () {
				// 用户确认分享后执行的回调函数
				},
				cancel: function () {}
			});
							
			// onMenuShareQZone
			wx.onMenuShareQZone({
		        title: data.title,
		        desc: data.desc,
		        link: data.link,
		        imgUrl: data.imgUrl,
				success: function () {},
				cancel: function () {}
			});
		})
	}	
			
	// 微信上传初始化
	this.wxupinit = function(callback) {
		
		// 判断是不是微信环境
		if (__ENV.env != 'weixin') return this;
		var api = 'chooseImage,previewImage,uploadImage';
		var weixin = 'http://res.wx.qq.com/open/js/jweixin-1.4.0.js';
		var config = '//config.suixingkang.com/js/weixin/?api=' + api;
		var media  = '//config.suixingkang.com/js/weixin/?media_id='
		$.getScript(weixin, function(){
			$.getScript(config, function(){});
		});
		return this;
	}
			
	// 获取微信图片
	this.wxImg = function(callback) {
		
		// 判断是不是微信环境
		if (__ENV.env != 'weixin') return this;

		wx.chooseImage({
			count: 1, // 默认9
			sizeType: ['original'],
			sourceType: ['album', 'camera'],
			success: function (res) {
			var localIds = res.localIds;
			// 上传图片
			wx.uploadImage({
				localId: localIds.toString(),
				isShowProgressTips: 1,
				success: function (res) {
				callback(res.serverId);
			  }
			});
		  }
		});
		return this;
	}
		
	// 微信上传
	this.upload = function(action, param, callback) {
		
		// 判断是不是微信环境
		if (__ENV.env != 'weixin') return this;
		
		// 使用URL上传
		this.wxImg(function(serverId) {
			param.code = 'wxres';
			param.mediaId = serverId;
			__ENV.json(action, param, callback);
		});
	}
	
	// 初始化分享
	this.share = function(data) {
		// 判断是不是微信环境
		this.shareData = data;
		if (__ENV.env != 'weixin') return this;
		
		var api = arguments[1] || 'onMenuShareTimeline,onMenuShareAppMessage,onMenuShareQQ,onMenuShareQZone';
		var weixin = 'http://res.wx.qq.com/open/js/jweixin-1.4.0.js';
		var config = '//config.suixingkang.com/js/weixin/?api=' + api;
		$.getScript(weixin, function(){
			$.getScript(config, function(){
				__this.wxRead(data);
			});
		});
		return this;
	};
	
	// QQ分享
	this.toQQ = function() {
		var data = arguments[0] || this.shareData;
		$("head").append('<meta itemprop="name"  content="'+data.title+'"/>');
		$("head").append('<meta itemprop="image" content="'+data.imgUrl+'"/>');
		$("head").append('<meta itemprop="share_url"  content="'+data.link+'"/>');
		$("head").append('<meta itemprop="description" itemprop="description" content="'+data.desc+'"/>');
		return this;
	}
	
	// 钉钉分享
	this.toDingTalk = function() {
		var data = arguments[0] || this.shareData;
		var dingJs = '//g.alicdn.com/ilw/ding/0.6.2/scripts/dingtalk.js';
		$.getScript(dingJs, function() {
			dd.ready(function () {
			      if (dd) {
			        dd.biz.navigation.setRight({
			          show: true,
			          control: true,
			          text: '',
			          onSuccess: function (result) {
			            dd.biz.util.share({
			              type: 0,
			              url: data.link,
			              title: data.title,
			              content: data.desc,
			              image: data.imgUrl,
			            })
			          }
			        })
			      }
			});
		});
	}
};// Import sxkui.js // 随行康前端框架 1.0
var MEDICAL_HEALTHUI = function(module) {
	
	// 全局配置
	var options = arguments[1] || {};
	this.config = $.extend(MEDICAL_HEALTH_CONF[module], options);
	this.config.module = module;
	
	
	// 运行环境判断
	var env = new MEDICAL_HEALTH_UI_ENV();
    if (env.isAndroid()) {
    	this.config.env = 'android';
		this.obj = new SXKUI_Android(this.config);
	}else if(env.isIOS()) {
		this.config.env = 'ios';
		this.obj = new SXKUI_IOS(this.config);
	} else if(env.isWeixin()) {
		this.config.env = 'weixin'
		this.obj = new SXKUI_Weixin(this.config);	
	} else {
		this.config.env = 'wap';
		this.obj = new SXKUI_WAP(this.config);	
	}
	
	return this.obj;
}
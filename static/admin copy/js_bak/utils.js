Array.prototype._indexOf = function(val, comparator) {
	for (var i = 0; i < this.length; i++) {
		if (comparator(this[i], val)) return i;
	}
	return -1;
};
// 删除数组中某个元素
Array.prototype._remove = function(val, comparator) {
	var index = this._indexOf(val, comparator);
	if (index > -1) {
		this.splice(index, 1);
	}
};
/**
 * 发送异步 http GET 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function httpGet(url, data, success, fail) {

	if (isObject(data)) { //传入了 data 参数
		url = buildQueryUrl(url, data);
	} else if(data) { // 没有传入 data 参数
		fail = success;
		success = data;
	}
	request(url, "GET", true, null, success, fail);
}

/**
 * 发送同步 http GET 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function asyncHttpGet(url, data, success, fail) {

	if (isObject(data)) { //传入了 data 参数
		url = buildQueryUrl(url, data);
	} else { // 没有传入 data 参数
		fail = success;
		success = data;
	}
	request(url, "GET", false, null, success, fail);
}

/**
 * 发送异步 http POST 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function httpPost(url, data, success, fail) {

	if (!isObject(data)) {
		fail = success;
		success = data;
		data = {};
	}
	request(url, "POST", true, data, success, fail);
}

/**
 * 发送 http 请求
 * @param url 地址
 * @param method 请求方式, POST, GET
 * @param async 是否同步
 * @param data 数据
 * @param success 成功时候回调
 * @param fail 失败时候回调
 * @returns {*}
 */
function request(url, method, async, data, success, fail) {

	if (typeof success != "function") {
		success = function() {}
	}
	if (typeof fail != "function") {
		fail = function() {}
	}
	var options = {
		type: method,
		url: url,
		async: async,
		data: data,
		success: function(result) {
			if (result.Code == "000") {
				success(result);
			} else {
				fail(result.Message);
			}
		},
		dataType: "json",
		error: function(error) {
			fail(error);
		}
	};
	$.ajax(options);
}

/**
 * 使用 json 传参
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function postJson(url, data, success, fail) {

	if (typeof success != "function") {
		success = function() {}
	}
	if (typeof fail != "function") {
		fail = function() {}
	}
	var options = {
		type: "POST",
		url: url,
		headers: {
			appId: "016726a0f24a0001a0",
			apiKey: "dfb28a624f0b4e2e2b45b1eddc81c3f0",
		},
		data: JSON.stringify(data),
		contentType: 'application/json;charset=utf-8',
		success: function(result) {
			if (result.Code == "000") {
				success(result);
			} else {
				fail(result.Message);
			}
		},
		dataType: "json",
		error: function(error) {
			fail(error);
		}
	};
	$.ajax(options);
}

/* 合并对象，使用 dist 覆盖 src */
function mergeObject(src, dist) {

	if (!isObject(src) || !isObject(dist)) {
		return;
	}
	for (var key in dist) {
		src[key] = dist[key];
	}
}

/* 判断是否是 javascript 对象 */
function isObject(obj) {
	return Object.prototype.toString.call(obj) == "[object Object]";
}

/* 判断是否是 javascript 对象 */
function isArray(arr) {
	return Object.prototype.toString.call(arr) == "[object Array]";
}

/* build query url */
function buildQueryUrl(url, params) {
	if (url.indexOf("?") == -1) {
		url += "?" + $.param(params);
	} else {
		url += "&" + $.param(params);
	}
	return url;
};

// 成功提示
function messageOk(message, callback, time) {
	//var layer = parent.layer === undefined ? layui.layer : top.layer;
	time = time ? time : 2000;
	JDialog.msg({
		type: "ok",
		content: message,
		timer: time,
		callback: callback
	});
}
// 失败提示
function messageError(message, callback) {
	JDialog.msg({
		type: "error",
		content: message,
		timer: 2000,
		callback: callback
	});
	return false;
}
function messageInfo(message, callback) {
	JDialog.msg({
		type: "info",
		content: message,
		timer: 1500,
		callback: callback
	});
	return false;
}
// 加载提示
function messageLoading(style) {
	if (style == 1) {
		return JDialog.msg({type:'loading', content: "正在加载中，请稍后...",timer:0, lock: true})
	} else {
		return JDialog.loader({
			timer : 0,
			lock : true,
			skin : 4
		});
	}
}

/**
 * 获取表单数据
 * @param formId
 * @returns {{}}
 */
function getFormData(formId) {
	var data = {};
	$.each($('#'+formId).serializeArray(), function(i,item) {
		data[item.name] = item.value;
	});
	return data;
}
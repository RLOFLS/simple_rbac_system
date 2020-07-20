/**
 * 地区选择器
 * @version 1.1.0
 * @author <yangjian102621@gmail.com>
 */
(function($) {

	$.fn.JAreaSelect = function(options) {

		var obj = {};
		var $container = $(this);
		var areaData = __AREADATA__;  //地区数据
		//初始化参数
		var defaults = {
			prov : 0, //省
			city : 0, //市
			dist : 0, //区
			level : 3, //选到第几级，默认第三级（地区）
			name : {
				prov : 'province',
				city : 'city',
				dist : 'dist'
			},
			onSelected : function (id, name) { //选中时回调

			},
			selectClassName : 'form-control', //select class名称

		};

		/* 合并参数 */
		options = $.extend(defaults, options);

		//创建元素
		obj.create = function() {

			obj.province = $('<select class="'+options.selectClassName+'" name="'+options.name.prov+'" required></select>');
			obj.province.append('<option value="0" selected>全部</option>');
			//加载所有省级
			$.each(areaData.prov, function(id, name) {
				if ( id == options.prov ) {
					obj.province.append('<option value="'+id+'" selected>'+name+'</option>');
				} else {
					obj.province.append('<option value="'+id+'">'+name+'</option>');
				}
			});

			//绑定选中省级事件
			obj.province.on('change', function() {

				//删除元素
				try {
					obj.city.remove();
					obj.dist.remove();
				} catch (e) {}

				var pid = $(this).val(); //获取省份id
				options.onSelected(pid, $(this).find("option:selected").text());

				if ( areaData.city[pid] && areaData.city[pid].length > 0 ) {

					obj.city = $('<select class="'+options.selectClassName+'" name="'+options.name.city+'" required></select>');
					obj.city.append('<option value="0" selected>全部</option>');
					$.each(areaData.city[pid], function(i, item) {
						if ( item.id == options.city ) {
							obj.city.append('<option value="'+item.id+'" selected>'+item.name+'</option>');
						} else {
							obj.city.append('<option value="'+item.id+'">'+item.name+'</option>');
						}
					});

					//切换城市的时候加载地区
					obj.city.on("change", function() {

						try {obj.dist.remove();} catch (e) {}
						//console.log(obj.getAreaString());

						var cid = $(this).val();
						options.onSelected(cid, $(this).find("option:selected").text());
						if ( options.level == 2 ) { //选择到城市一级
							return false;
						}

						if ( areaData.dist[cid] && areaData.dist[cid].length > 0 ) {
							obj.dist = $('<select class="'+options.selectClassName+'" name="'+options.name.dist+'" required></select>');
							obj.dist.append('<option value="0" selected>全部</option>');
							$.each(areaData.dist[cid], function(i, item) {
								if ( item.id == options.dist ) {
									obj.dist.append('<option value="'+item.id+'" selected>'+item.name+'</option>');
								} else {
									obj.dist.append('<option value="'+item.id+'">'+item.name+'</option>');
								}
							});
							obj.dist.on("change", function () {
								options.onSelected($(this).val(), $(this).find("option:selected").text());
							});
							$container.append(obj.dist);

						}
					});
					$container.append(obj.city);
					obj.city.trigger("change"); //自动触发事件

				}

			});
			$container.append(obj.province);
			if ( options.level == 1 ) { //选择到省一级
				obj.province.off("change");
			}
			obj.province.trigger("change"); //自动触发事件
		}

		//获取区域id
		obj.getAreaId = function() {
			return[
				obj.province.val(),
				obj.city ? obj.city.val() : 0,
				obj.dist ? obj.dist.val() : 0
			];
		}

		//获取省份名称
		obj.getProvince = function() {
			return obj.province.find("option:selected").html();
		}

		//获取城市名称
		obj.getCity = function() {
			try {
				return obj.city.find("option:selected").html();
			} catch (e) {
				return "";
			}
		}

		//获取地区名称
		obj.getDist = function() {
			try {
				return obj.dist.find("option:selected").html();
			} catch (e) {
				return "";
			}
		}

		//获取区域字符串
		obj.getAreaString = function() {
			var html = [];
			html.push(obj.province.find("option:selected").html());
			try {
				html.push(obj.city.find("option:selected").html());
				html.push(obj.dist.find("option:selected").html());
			} catch (e) {}
			return html;
		}

		obj.create();
		return obj;

	}

	$.extend({
		getAreaById : function (prov, city, dist) {
			var s = [];
			for (var o in __AREADATA__["prov"]) {
				if (prov == o) {
					s.push(__AREADATA__["prov"][o]);
				}
			}
			if(city) {
				for (var o in __AREADATA__["city"]) {
					for (var k in __AREADATA__["city"][o]) {
						if (__AREADATA__["city"][o][k]['id'] == city) {
							s.push(__AREADATA__["city"][o][k]["name"]);
						}
					}
				}
			}
			if(dist) {
				for (var o in __AREADATA__["dist"]) {
					for (var k in __AREADATA__["dist"][o]) {
						if (__AREADATA__["dist"][o][k]['id'] == dist) {
							s.push(__AREADATA__["dist"][o][k]["name"]);
						}
					}
				}
			}
			return s;
		}
	})
})(jQuery);

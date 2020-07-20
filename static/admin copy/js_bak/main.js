/**
 * @author yangjian
 * @since 18-12-25 下午2:47.
 */

$(document).ready(function() {

	// 初始化全选事件
	$('#select-all').change(function() {
		if (this.checked) {
			$('#data-list-form input[name="ids[]"]').prop("checked", true);
		} else {
			$('#data-list-form input[name="ids[]"]').prop("checked", false);
		}
	});

	$.fn.extend({
		// add button loading extension
		jbutton: function(flag) {
			if (flag == "loading") {
				$(this).addClass("btn-disabled");
				$(this).prop("disabled", true);
				var text = $(this).text();
				$(this).data("text", text);
				var loadingText = $(this).attr("data-loading-text") || "正在提交...";
				$(this).text(loadingText);
			} else if(flag == "reset") {
				var text = $(this).data("text");
				$(this).prop("disabled", false);
				$(this).removeClass("btn-disabled");
				$(this).text(text);
			}
		},
	});
});

// 删除一条数据
function removeItem(url, item, callback) {

	JDialog.confirm({
		title : "删除提示",
		width : 400,
		lock : true,
		effect : 1,
		maxEnable : false,
		content : '确定要删除该条记录吗？',
		icon : 'warn',
		offset : 'cc',
		effect : 'zoomIn',
		button : {
			'确认' : function(dialog) {
				dialog.close();
				var loading = messageLoading(1);
				setTimeout(function() {
					httpGet(url, {id: item.id}, function() {
						callback();
						loading.hide();
						messageOk("删除成功.");
					}, function() {
						loading.hide();
						messageError("删除失败.")
					});
				}, 500);
			},
			'取消' : function(dialog) {
				dialog.close();
			}
		}
	});

}
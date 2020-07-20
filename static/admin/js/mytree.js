var myTree = {
    "data" : {
        'open-icon' : 'ace-icon tree-minus',
        'close-icon' : 'ace-icon tree-plus',
    },
    "fopen" : function ($obj) { //有儿子打开
        var that = this;
        that.fchangeOpenStatus($obj);
        //console.log($obj.find("li"));
        $obj.find("li").each(function (idx,item) {
            that.fopen($(item));
        })
    },
    "fchangeOpenStatus" : function ($obj) {
        var that = this;
        $obj.addClass("tree-selected");
        $obj.children().eq(2).removeClass("hidden");
        $obj.find("> .tree-branch-header .icon-folder").eq(0).removeClass(that.data['close-icon']).addClass(that.data['open-icon']);
        $obj.attr("aria-expanded", "true");
        $obj.addClass("tree-open");
    },
    "fclose" : function($obj) { //有儿子关闭
        var that = this;
        $obj.removeClass("tree-selected");
        $obj.children().eq(2).addClass("hidden");
        $obj.find("> .tree-branch-header .icon-folder").eq(0).removeClass(that.data['open-icon']).addClass(that.data['close-icon']);
        $obj.attr("aria-expanded", "false");
        $obj.removeClass("tree-open");
        $obj.find("li").each(function (idx,item) {
            that.fclose($(item));
        })
    },
    "close":function($obj){//无儿子关闭
        $obj.removeClass("tree-selected");
    },
    "open":function($obj){ //有儿子开启
        $obj.addClass("tree-selected");
    },
    "isFolder":function($obj){
        return $obj.hasClass("tree-branch")?true:false;
    },
    "isSelect":function($obj){
        return $obj.hasClass("tree-selected")?true:false;
    },

}
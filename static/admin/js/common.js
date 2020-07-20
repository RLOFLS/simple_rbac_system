toastr.options = { // toastr配置
    "showDuration": "200", //显示的动画时间
    "hideDuration": "1000", //消失的动画时间
    "timeOut": "2000", //展现时间
    "extendedTimeOut": "1000", //加长展示时间

}

//菜单
if ($(".nav-list .active").is("li")) {
    if($(".nav-list .active").parent().parent().is("li")){
        $(".nav-list .active").parent().parent().attr("class","open");
        $(".nav-list .active").parent().css('display','block');
    }
}

function logout(){
    $.ajax({
        type : "POST",
        url : "/admin/login/logout",
        data : '',
        dataType: "json",
        //请求成功
        success : function(result) {
            console.log(result);
            if ( result.Code == 000 ){
                alert(result.Message);
                location.href = '/';
            }
            
            else
            alert( result.Message);
        },
        //请求失败，包含具体的错误信息
        error : function(e){
            alert('网络繁忙');
        }
    });
}

 //修改密码
 function modifyPassword(){
    $("#modifyPasswordModal").modal("show");
}
//二次确认
function confirmModifyPassword(){
    var newPassword = $("#newPassword").val();
    if(newPassword == '') {
        toastr.success("请输入新密码")
        return false;
    }

    $.ajax({
        type : "POST",
        url : "/admin/jump/modifypassword",
        data : { "newPassword" : newPassword },
        dataType: "json",
        //请求成功
        success : function(result) {

            $("#modifyPasswordModal").modal("hide");
            if ( result.Code == 000 ){
                alert("修改成功，稍后将退出")

                setTimeout(function () {
                    logout()
                },3000);	

            }else{
                toastr.error(result.Message);
            }
        },
        //请求失败，包含具体的错误信息
        error : function(e){
            $("#modifyPasswordModal").modal("hide");
            toastr.error("系统繁忙，请稍候再试");
        }
    });

}

 //权限按钮删除
 function deleteNoPermissionButton() {
    if (JSON.stringify(noOperations)=="{}") {
        return
    }
    noOperations.forEach(item => {
        let btn_class = '.permission-'+item ;
        $(btn_class).remove();
    });
}



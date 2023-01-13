
$(document).ready(function () {
    let cid=GetQueryString("cid")
    $.ajax({
        type:"get",
        url:"/product/selByCid?cid="+cid,
        dataType:'json',
        success:function (res){
            let obj=eval(res.list)
            console.log(obj)
            let str="<ul class=\"cate_list\">\n"
            for(let i=0;i<obj.length;i++){
                str+="                \t<li>\n" +
                    "                    \t<div class=\"img\"><a href=\"/product/info?id="+obj[i].id+"\"><img src=\""+obj[i].imgPath+obj[i].img+"\" width=\"210\" height=\"185\" /></a></div>\n" +
                    "                        <div class=\"price\">\n" +
                    "                            <font>￥<span>"+obj[i].price+"</span></font> &nbsp; 26R\n" +
                    "                        </div>\n" +
                    "                        <div class=\"name\"><a href=\"#\">"+obj[i].name+"</a></div>\n" +
                    "                        <div class=\"carbg\">\n" +
                    "                        \t<a href=\"#\" class=\"ss\">收藏</a>\n" +
                    "                            <a href=\"#\" class=\"j_car\">加入购物车</a>\n" +
                    "                        </div>\n" +
                    "                    </li>"
            }
            str+="</ul>"
            str+="<div class=\"pages\">\n" +
                "                \t<a href=\"#\" class=\"p_pre\">上一页</a><a href=\"#\" class=\"cur\">1</a><a href=\"#\">2</a>" +
                "<a href=\"#\">3</a>...<a href=\"#\">20</a><a href=\"#\" class=\"p_pre\">下一页</a>\n" +
                "                </div>"
            $('#plist_cid').empty()
            $('#plist_cid').append(str)
        },
    })
})


function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg); //获取url中"?"符后的字符串并正则匹配
    var context = "";
    if (r != null)
        context = decodeURIComponent(r[2]);
    reg = null;
    r = null;
    return context == null || context == "" || context == "undefined" ? "" : context;
}
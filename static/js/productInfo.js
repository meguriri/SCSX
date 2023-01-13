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
var pid
$(document).ready(function () {
    pid=GetQueryString("id")
    console.log("pid:",pid)
    $.ajax({
        type:"get",
        url:"/product/productDetail?id="+pid,
        dataType:'json',
        success:function (res){
            let obj=JSON.parse(res.list)
            console.log(obj)
            $('#p_name').html("<p>"+obj.name+"</p>")
            $('#p_price').html("￥"+obj.price)
            $('#p_type').html(obj.type)
            $('#p_color').html(obj.color)
            $('#p_img').attr("src",obj.imgPath+obj.img)
            $('#tsImgS').children('a').attr("href",obj.imgPath+obj.img)
        },
    })
    $('#addUpdate').click(function (){
        var c = $(".n_ipt").val();
        c=parseInt(c)+1;
        $(".n_ipt").val(c);
    })
    $('#jianUpdate').click(function (){
        var c = $(".n_ipt").val();
        if(c==1){
            c=1;
        }else{
            c=parseInt(c)-1;
            $(".n_ipt").val(c);
        }
    })


    $('#showDiv').click(function (){
        let n=$('#num').val()
        $.ajax({
            type: "post",
            url: "/product/addCart" ,
            dataType: 'json',
            data:{"pid":pid,"num":n},
            success: function (res) {
                if(res.msg=="fail"){
                    alert("请登录")
                    window.location.href="/user/login"
                }else {
                    console.log(res.type,res.total,res.price)
                    $('#MyDiv1').attr("style", "display:block")
                    let bgdiv = $('#fade1')
                    bgdiv.attr("style", "display:block")
                    $('cartmsg').empty()
                    $('#cartmsg').append("<span style=\"color:#3e3e3e; font-size:18px; font-weight:bold;\">" +
                        "宝贝已成功添加到购物车</span><br />\n" +
                        "\t购物车共有"+res.type+"种宝贝（"+res.total+"件） &nbsp; &nbsp; 合计："+res.price+"元")
                    bgdiv.style.width = document.body.scrollWidth;
                    $("#fade1").height($(document).height())
                }
            }
        })
    })

    $('#closeDiv').click(function (){
        $('#MyDiv1').attr("style","display:none")
        $('#fade1').attr("style","display:none")
    })
})
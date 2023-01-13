var id
$(document).ready(function (){
    $.ajax({
        type:"get",
        url:"/buyCar/allProducts",
        dataType:'json',
        success:function (res){
            if(res.msg=="fail"){
                alert("请登录")
                window.location.href="/user/login"
            }else {
                console.log("list:",res.list)
                console.log("total:",res.total)
                let obj=eval(res.list)
                console.log("obj:",obj)
                let str=''
                for(let i=0;i<obj.length;i++){
                    str+='<tr class="car_tr">\n' +
                        '            <td>\n' +
                        '            \t<div class="c_s_img"><img src="'+obj[i].imgPath+'" width="73" height="73" /></div>\n' +
                        obj[i].name +
                        '            </td>\n' +
                        '            <td align="center">颜色：'+obj[i].color+'</td>\n' +
                        '            <td align="center">\n' +
                        '            \t<div class="c_num">\n' +
                        '                   <input type="text" value="'+obj[i].number+'" name="" class="car_ipt" disabled />  \n' +
                        '                </div>\n' +
                        '            </td>\n' +
                        '            <td align="center" style="color:#ff4e00;">￥'+obj[i].price+'</td>\n' +
                        '            <td align="center"><a onclick="showDiv()" class="del" id="'+obj[i].id+'">删除</a></td>\n' +
                        '          </tr>'
                }
                $('#title').after(str)
                $('#totalPrice').html('￥'+res.total)
            }
        },
    })

    $('#closeDiv,#cancel').click(function (){
        $('#MyDiv').attr("style","display:none")
        $('#fade').attr("style","display:none")
    })

    $('#sure').click(function (){
        $.ajax({
            type:"post",
            url:"/buyCar/delBuyCar",
            data: {"id":id},
            dataType:'json',
            success:function (res){
                if(res.msg=="fail"){
                    alert("删除失败")
                }else {
                    location.reload();
                }
            },
        })
    })

    $('#tab').on('click', '.del', function () {
        id=$(this).attr('id')
        console.log('id:',id)
    });
})

function showDiv(){
    $('#MyDiv').attr("style", "display:block")
    let bgdiv = $('#fade')
    bgdiv.attr("style", "display:block")
    bgdiv.style.width = document.body.scrollWidth;
    $("#fade").height($(document).height())
}

function clearMyCar(){
    if($('#clear').is(':checked')){
        if(confirm("确定清空购物车嘛?")){
            $.ajax({
                type:"post",
                url:"/buyCar/delAllBuyCar",
                dataType:'json',
                success:function (res){
                    if(res.msg=="fail"){
                        alert("删除失败")
                    }else {
                        alert("删除成功")
                        location.reload();
                    }
                },
            })
        }
    }
}
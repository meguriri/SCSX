$(document).ready(function () {
    $.ajax({
        type:"get",
        url:"/order/allProducts",
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
                        '                <td>\n' +
                        '                    <div class="c_s_img"><img src="'+obj[i].imgPath+'" width="73" height="73" /></div>\n' +
                        obj[i].name +
                        '                </td>\n' +
                        '                <td align="center">颜色：'+obj[i].color+'</td>\n' +
                        '                <td align="center">'+obj[i].number+'</td>\n' +
                        '                <td align="center" style="color:#ff4e00;">￥'+obj[i].price+'</td>\n' +
                        '              </tr>'
                }
                $('#title').after(str)
                $('#totalPrice').html('商品总价：￥'+res.total)
                $('#yfk').html('￥'+res.total)
            }
        },
    })

    $('#confirmOrder').click(function (){
        let person=$('#person').val()
        let tel =$('#tel').val()
        let address=$('#address').val()
        if(person.trim()===''){
            alert('请填写收货人姓名')
            $('#username').focus()
            return
        }
        if(tel.trim()===''){
            alert('请填写联系方式')
            $('#username').focus()
            return
        }
        if(address.trim()===''){
            alert('请填写收货地址')
            $('#username').focus()
            return
        }else{
            $.ajax({
                type:"post",
                url:"/order/confirmOrder",
                dataType:'json',
                data:{"person":person,"tel":tel,"address":address},
                success:function (res){
                    if(res.msg=="ok"){
                        console.log(res.num)
                        window.location.href="/order/okOrder?oid="+res.num
                    }
                }
            })
        }
    })
})
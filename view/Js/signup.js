function SignUp(){
    var data={
        fname:document.getElementById("fname").value,
        lname:document.getElementById("lname").value,
        pw:document.getElementById("password").value,
        email:document.getElementById("email").value,
        role:document.getElementById("role").value

    }
    let conf_pw=document.getElementById("conf_pw").value

    if (data.pw!=conf_pw){
        alert("Password doesn't match")
        return
    }

    fetch('/signup',{
        method:"POST",
        body:JSON.stringify(data),
        headers:{"Content-type":"application/json"}
    })
    .then(response=>{
        if(response.ok){
            window.open('index.html','_self')
        }else{
            throw new Error(response.statusText)
        }
    })
    .catch(e=>{
        alert(e.message)
        return
    })
}
function login(){
    let data={
        email:document.getElementById("email").value,
        pw:document.getElementById("password").value
    }

    fetch('/login',{
        method:"POST",
        body:JSON.stringify(data),
        headers:{"Content-type":"application/json"}
    })
    .then(response=>{
        if (response.ok){
            window.open('Explore.html','_self')
        }else{
            throw new Error(response.statusText)
        }
    })
    .catch(e =>{
        if (e.message=="Unauthorized"){
            alert(e.message + "-Credintials does not match")
            return
        }else if(e.message=="Not Found"){
            alert(e.message + "-User with the given email not found")
            return
        }else{
            alert(e.message)
            return
        }
    })
}

function logout(){
    fetch('/logout',)
    .then(response=>{
        if(response.ok){
            window.open("index.html",'_self')
        }else{
            throw new Error(response.statusText)
        }
    })
    .catch(e=>{
        alert(e.message)
        return
    })
}
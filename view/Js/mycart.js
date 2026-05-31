function GetUserMail(){
    return fetch('/getsession',{
        method:"GET",
    })
    .then(response=>{
        if(response.ok){
            return response.json() //we want only response body parse in JSON to have better access...
            
        }else{
            throw new Error(response.statusText)
        }
    })
    .then(data=>{
        return data["Email"] //return value email...
    })
    .catch(e=>{
        alert(e.message)
        throw e
    })
}

document.addEventListener('DOMContentLoaded',function(){
    GetUserMail()
    .then(user_email=>{
        if (user_email){
                return fetch(`/readAllPurchases/${user_email}`,{
        method:"GET"
    })
    }else{
            throw new Error("Please Log In first")
        }
    })
    .then(response=>{
        if(response.ok){
           return response.json()
        }else{
            throw new Error(response.statusText)
            return
        }
    })
    .then(courses=>{
        displayPurchase(courses)

    })
    .catch(e=>{
        alert(e.message)
        return
    })
})



function displayPurchase(courses){
    let container=document.getElementById("Purchased_container")

    courses.forEach(course => {
        let card = `<div class="col">
  <div class="card border-0 rounded-3" onclick='ShowModal(${JSON.stringify(course)})'>
    <img src="${course.image_file}" class="card-img-top" alt="Product Image"/>
    <div class="card-body bg-black text-white">
      <h6 class="card-title fw-bold">${course.title}</h6>
      <div class="d-flex align-items-center gap-2 mb-1">
        <p>Author: <span class="small">${course.author_name}</span></p>
      </div>
    </div>
    <div class="card-footer bg-black border-0 pb-3">
      <span class="badge rounded-pill px-3 py-2" style="background:#ff90e8; color:#000;">
        Purchased
      </span>
      <button onclick="CopyLink(${course.c_id})" 
              class="btn btn-outline-light rounded-pill py-2 fw-semibold w-100 mt-2">
        🔗 Copy Share Link
      </button>
    </div>
  </div>
</div>`
  container.innerHTML+=card
        
    });

}

let selected_card=null

function ShowModal(course){
    selected_card=course
    document.getElementById("Course_Image").src=course.image_file
    document.getElementById("Course_Title").innerText=course.title
    document.getElementById("Course_Author").innerText=course.author_name
    document.getElementById("Course_Description").innerText=course.description
    document.getElementById("open_pdf").href=course.pdf_file

    let purchasedmodal=document.getElementById("Purchased_Modal")
    let modal=new bootstrap.Modal(purchasedmodal)

    modal.show()

}

function CopyLink(c_id){
    let link=`${window.location.origin}/Explore.html?id=${c_id}`
    if (navigator.clipboard){
        navigator.clipboard.writeText(link)
        alert("Link Copied")
    }else{
        prompt("copy link here",link)
    }
}

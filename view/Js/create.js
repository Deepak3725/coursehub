//Opening Uploading interface for user
let modal;
function uploadModal(){

    document.getElementById("upload_header").innerText="Upload Your Course"//because its changed during edit modal
    let submit_btn=document.getElementById("submit_btn")
    let upload_Modal=document.getElementById("Upload_modal")
    modal=new bootstrap.Modal(upload_Modal)

    modal.show()

    submit_btn.onclick = function(){
        GetUserMail().then(creator_mail => CreateCourse(creator_mail))
    }

}
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



function CreateCourse(user_mail){
    //Getting course data..

let formdata=new FormData() //create empty form container to send data to backend
formdata.append("title",document.getElementById("title").value)//add in container in key value pair..
formdata.append("author_name",document.getElementById("author_name").value)
formdata.append("price",document.getElementById("price").value)
formdata.append("description",document.getElementById("description").value)
formdata.append("image_file",document.getElementById("image_file").files[0])
formdata.append("pdf_file",document.getElementById("pdf_file").files[0])//sends first selected files...
formdata.append("creator_mail",user_mail)
formdata.append("bank_name",document.getElementById("bank_name").value)
formdata.append("acc_num",document.getElementById("acc_num").value)
formdata.append("bank_qr",document.getElementById("bank_qr").files[0])
fetch('/addcourse',{
    method:"POST",
    body:formdata //my request body will take form to backend

    //No header needed coz browser will set it auto to form...
})
.then(response=>{
    if (response.ok){
        modal.hide()
    }else{
        throw new Error(response.statusText)
    }
})
.catch(e=>{
    alert(e.message)
    return
})

}

function CopyLink(c_id){
    let link=`${window.location.origin}/Explore.html?id=${c_id}`//window.location.orgin:saves base URL:http:localhost/8080
    //Navigator clipboard use to copy link in our clipboards it has some restrictions with http
    if(navigator.clipboard){ //if no restriction
        navigator.clipboard.writeText(link)//link copied to clipboard
        alert("Link copied to your clipboard")
    }else{
        prompt("copy this link",link)//if restriction shows prompt where we can manually copy..
    }
}

document.addEventListener('DOMContentLoaded',function(){
    GetUserMail()
    .then(creator_mail=>{
        if(creator_mail){
            return fetch(`/getsellincourse/${creator_mail}`,{method:"GET"})
        }else{
            throw new Error("Please Log In First")
        }
    })
    .then(response=>{
        if(response.ok){
            return response.json()
        }else{
            throw new Error(response.statusText)
        }
    })
    .then(Selling_list=>{
        displaySelling(Selling_list)
    })
    .catch(e=>{
        alert(e.message)
    })
})

function displaySelling(Sell_Course){
    let container=document.getElementById("sell_container")

    Sell_Course.forEach(Course=> {
        let card=` <div class="col">
    <div class="card border-0 rounded-3 overflow-hidden" onclick='showEditModal(${JSON.stringify(Course)})'>
      <img src="${Course.image_file}" class="card-img-top" alt="Product Image"/>
      <div class="card-body bg-black text-white">
        <h6 class="card-title fw-bold">${Course.title}</h6>
        <div class="d-flex align-items-center gap-2 mb-1">
          <p>Author: <span class="small">${Course.author_name}</span></p>
        </div>
      </div>
      <div class="card-footer bg-black border-0 pb-3">
        <span class="badge rounded-pill px-3 py-2" style="background:#ff90e8; color:#000;">
          Nu ${Course.price}
        </span>
        <button onclick="event.stopPropagation();CopyLink(${Course.id})" 
                    class="btn btn-outline-light rounded-pill py-2 fw-semibold w-100 mt-2">
                🔗 Copy Share Link
            </button>
      </div>
    </div>
  </div>
        `//event.stopPropagation to avoid clicking enter card...
      container.innerHTML+=card   
    });
     
}


//Updating the card:

let selected_card=null

function showEditModal(course){
    selected_card=course

    let update_btn=document.getElementById("submit_btn")
    update_btn.innerText="Update Course"
    document.getElementById("upload_header").innerText="Update Your Course Here"
    document.getElementById("title").value=course.title
    document.getElementById("author_name").value=course.author_name
    document.getElementById("price").value=course.price
    document.getElementById("description").value=course.description
    document.getElementById("bank_name").value=course.bank_name
    document.getElementById("acc_num").value=course.acc_num

   //all the fields here will take values already and show you empty image and pdf if you make upload it will change if not remains same

   let update_Modal=document.getElementById("Upload_modal")
   modal=new bootstrap.Modal(update_Modal)

   modal.show()

   update_btn.onclick=function(){
    UpdateCourse(selected_card)
   }

}

function UpdateCourse(selected_card){
    let formdata=new FormData()
    formdata.append("title",document.getElementById("title").value)//add in container in key value pair..
    formdata.append("author_name",document.getElementById("author_name").value)
    formdata.append("price",document.getElementById("price").value)
    formdata.append("description",document.getElementById("description").value)
    formdata.append("creator_mail",selected_card.creator_mail)
    formdata.append("bank_name",document.getElementById("bank_name").value)
    formdata.append("acc_num",document.getElementById("acc_num").value)

    let imgFile=document.getElementById("image_file")

   if (imgFile.files.length>0){ //check if image file is empty
    formdata.append("image_file",imgFile.files[0])//not empty append the image file
   }
   let pdfFile=document.getElementById("pdf_file")
   if (pdfFile.files.length>0){
    formdata.append("pdf_file",pdfFile.files[0])
   }
   let bank_qr=document.getElementById("bank_qr")

   if (bank_qr.files.length>0){
    formdata.append("bank_qr",bank_qr.files[0])
   }

    fetch(`/updatecourse/${selected_card.id}`,{
        method:"PUT",
        body:formdata
    })
    .then(response=>{
        if (response.ok){
            alert("Course Update Successful")
            modal.hide()
        }else{
            throw new Error(response.statusText)
            return
        }
    })
    .catch(e=>{
        alert(e.message)
    })

}
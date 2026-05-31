function GetUserMail() {
    return fetch('/getsession', {
        method: "GET",
    })
    .then(response => {
        if (response.ok) {
            return response.json()
        } else {
            throw new Error(response.statusText)
        }
    })
    .then(data => {
        return data["Email"]
    })
    .catch(e => {
        alert(e.message)
        throw e  // Re-throw so outer catch can handle it
    })
}
document.addEventListener("DOMContentLoaded",function(){
    fetch('/getallcourses',{
        method:"GET"
    })
    .then(response=>{
        if (response.ok) {
            return response.json() //converts whatever response I get in json so, next I will get my actual data
            
        }else{
            throw new Error(response.statusText)
        }
    })
    
    //now I will get array of courses that my backend has returned.. that is courses here
    .then(courses=>{
        let container=document.getElementById("card_container") 
        if (!courses || courses.length === 0) {
        container.innerHTML = "<p class='text-white text-center'>No courses available yet!</p>"
        return
    }
        displayCourses(courses)
        let param=new URLSearchParams(window.location.search) //stores url in the param
        let id=param.get("id") //get id from url

        if (id){
            let course=courses.find(c=>c.id==id) // we find course from courses array having course_id =id
            if(course){
                showModal(course)//if course found then show its modal
            }
        }


    })
    .catch(e=>{
        alert(e.message)
        return
    })
})

function displayCourses(courses){
    let container=document.getElementById("card_container") //gets empty card later we add it
    courses.forEach(course => {//Now we just create individual cards using same html card
        let card=` <div class="col">
    <div class="card border-0 rounded-3 overflow-hidden" onclick='showModal(${JSON.stringify(course)})'>
      <img src="${course.image_file}" class="card-img-top" alt="Product Image"/>
      <div class="card-body bg-black text-white">
        <h6 class="card-title fw-bold">${course.title}</h6>
        <div class="d-flex align-items-center gap-2 mb-1">
          <p>Author: <span class="small">${course.author_name}</span></p>
        </div>
      </div>
      <div class="card-footer bg-black border-0 pb-3">
        <span class="badge rounded-pill px-3 py-2" style="background:#ff90e8; color:#000;">
          Nu ${course.price}
        </span>
      </div>
    </div>
  </div>
        `
   container.innerHTML+=card //Inside container(HTML) write this html code...
        
    });
}

let selected_card=null

function showModal(course){
    selected_card=course
    document.getElementById("Course_Image").src=course.image_file
    document.getElementById("Course_Title").innerText=course.title
    document.getElementById("Course_Author").innerText=course.author_name
    document.getElementById("Course_Description").innerText=course.description
    document.getElementById("Course_Price").innerText=course.price

    let coursemodal=document.getElementById("Course_Modal")

    let modal=new bootstrap.Modal(coursemodal)

    modal.show()
}

function showPaymentModal(course){
    console.log(course)
    document.getElementById("bank_qr").src=course.bank_qr
    document.getElementById("BankName").innerText=course.bank_name
    document.getElementById("AccNum").innerText=course.acc_num
    bootstrap.Modal.getInstance(document.getElementById("Course_Modal")).hide();
    let paymentModal=document.getElementById("Payment_Modal")

    let modal=new bootstrap.Modal(paymentModal)
    modal.show()
    
}
function ConfirmPayment(selected_card){
    let reference_num = document.getElementById("reference_num").value

    if(reference_num.trim() == ""){  //trim() removes white spaces.. sometimes one may just enter space bar...
        alert("Enter JRNL Number, it's necessary")
        return
    }

    GetUserMail()
    .then(email => {
        let data = {
            buyer_mail: email,
            course_id: selected_card.id,
            ref_num: reference_num,
        }
        return fetch('/initiatepayment', {
            method: "POST",
            body: JSON.stringify(data),
            headers: {"Content-Type": "application/json"}
        })
    })
    .then(response => {
        if(response.ok){
            alert("Payment request sent! Creator will verify and you'll get a confirmation email.")
        }else{
            throw new Error(response.statusText)
        }
    })
    .catch(e => {
        alert(e.message)
    })
}
function CopyLink(selected_card){
    let link=`${window.location.origin}/Explore.html?id=${selected_card.id}`

    if (navigator.clipboard){
        navigator.clipboard.writeText(link)
        alert("Link copied to your clipboard")
    }else{
        prompt("Copy this link",link)
    }
}
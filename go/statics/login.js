// Global variable
var email_RegEx = /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/;
var password_RegEx = /^(?=.*[a-z])(?=.*[A-Z])^(?=.*[0-9])(?=.*[.,:;!?#@&_\-])[a-zA-Z0-9.,:;!?#@&_\-]{8,}$/;

// Creating the URL (login server endpoint) variable.
var baseURL = "http://localhost:8080";
const ApiEndpoint = "/api/v1"

async function SignIn_SendData (emailValue, passwordValue) {
	if (email_RegEx.test(emailValue) === true){
    	let response = await request(ApiEndpoint + "/signin","POST", {email: emailValue, password: passwordValue});
        console.log(response);
        if (response.status != 200) {
            alert(response.error + "\n\n" + response.solution);
        } else {
        	location.replace(response.action);
        }
	} else {
    	alert("The conditions of the RegEx were not respected:\n\nMake sure you write an Username with the following syntax:\n<Surname>.<Name>.<SubscriptionYear2Digit><S|T><SerialNumber>\n\nExample:\nCognome.Nome.22S000\n");
    }
}

async function SignUp_SendData (emailValue, passwordValue, passwordCheckValue) {
	//console.log("SignUp | ? | {", usernameValue, ", ", passwordValue, ", ", passwordCheckValue, "};");
    if (email_RegEx.test(emailValue) === true){
		if (password_RegEx.test(passwordValue) === true) {
        	if (passwordValue === passwordCheckValue) {
            	//console.log("SignUp | v | {", usernameValue, ", ", passwordValue, ", ", passwordCheckValue, "};");
        		let response = await request(ApiEndpoint + "/signup","POST", {email: emailValue, password: passwordValue});
				console.log(response);
                if (response.status != 200) {
                    alert(response.error + "\n\n" + response.solution);
                } else {
                	switchTo("SignIn");
                }
            } else {
            	alert("Password and ConfirmPassword do not contain the same values.");
            }
        } else{
        	alert("The conditions of the RegEx were not respected:\n\nMake sure you write a Password with the following characteristics:\n - At least 8 characters\n - At least 1 uppercase letter\n - At least 1 lowercase letter\n - At least 1 digit\n - At least 1 of this special character: .,:;!?#@&_\-\n\nExample:\nCognomeNome_22S000\n");
        }
        
        
	} else {
    	alert("The conditions of the RegEx were not respected:\n\nMake sure you write an Username with the following syntax:\n<Surname>.<Name>.<SubscriptionYear2Digit><S|T><SerialNumber>\n\nExample:\nCognome.Nome.22S000\n");
    }
}

async function request (endpoint, requestMethod, requestBody) {
    let response = await fetch(baseURL + endpoint, {
        method: requestMethod,
        header: {
            "Content-Type": "application/json",
			"charset": "UTF-8"
        },
        body: JSON.stringify(requestBody)
    });

    let returnValueText = await response.json();
    return returnValueText;
}

/* Valid user for example login:
	let users = [
    	{username: "Cognome.Nome.22S000", password: "CognomeNome_22S000"}, // Student Login
    	{username: "Cognome.Nome.22T000", password: "CognomeNome_22T000"}  // Teacher Login
    ];
*/
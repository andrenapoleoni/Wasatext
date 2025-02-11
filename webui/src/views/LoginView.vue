<script>
export default {
    data() {
        return {
            username: "",
            errorMsg: "",
            usernameValidation: new RegExp('^\\w{3,16}$'),
        }
    },
    emits: ['login-success'],
    methods: {
        async doLogin() {
            try {
                if (this.username.length < 3 || this.username.length > 16) throw "Invalid username, it must contains min 3 characters and max 16 characters"
                if (!this.usernameValidation.test(this.username)) throw "Invalid username, it must contain only letters and numbers"
                let response = await this.$axios.post('/session', {
                    username: this.username,
                });
                sessionStorage.userID = response.data.user.userID;
                sessionStorage.username = response.data.user.username;
                sessionStorage.token = response.data.token;
                sessionStorage.photo = response.data.user.photo;
                window.location.reload();   
                
                this.$router.push("/home");
                this.$emit('login-success');
            } catch (e) {
                this.errorMsg = e.toString();
                document.getElementsByTagName("input")[0].style.outline = "auto";
                document.getElementsByTagName("input")[0].style.outlineColor = "red";
            };
        }
    },
    mounted() {
        if (sessionStorage.token) {
            this.$router.push("/home");
            return;
        }







        sessionStorage.clear();
    },
}

</script>

<template>
    <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
    <div class="login-container">
        <form @submit.prevent="doLogin">
            <h1>WasaText</h1>
            <input type="text" v-model="username" placeholder="Enter your username" />
            <button type="submit">Login</button>
        </form>
    </div>
</template>


<style>
.login-container {
display: flex;
justify-content: center;
align-items: center;
height: 50vh;
}

.login-container form {
display: flex;
flex-direction: column;
align-items: center;
}

.login-container input {
margin: 15px;
padding: 10px;
border: 1px solid #ccc;
border-radius: 5px;
}

.login-container button {
padding: 10px 20px;
border: none;
border-radius: 5px;
background-color: #007bff;
color: white;
cursor: pointer;
}

.login-container button:hover {
background-color: #0056b3;
}
</style>
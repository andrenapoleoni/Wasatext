<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {
  data() {
    return {
      user: {
		userID: sessionStorage.userID,
        username: sessionStorage.username,
        photo: sessionStorage.photo, 
      },

	  isLoggedIn: sessionStorage.token ? true : false,
	  showUserBanner: false
    };
  },
  methods: {
	logout() {
	  sessionStorage.clear();
	  this.$router.push("/");
	},
	handleLoginSuccess() {
      this.isLoggedIn = true;
      window.location.reload();
    },
	goToProfile() {
		localStorage.username = this.user.username;
		localStorage.photo = this.user.photo;
		localStorage.userID = this.user.userID;
		console.log(this.user,this.username);
		this.showUserBanner=false;
		
	  this.$router.push('/profile');
	},
	goHome() {
      this.$router.push('/home').then(() => {
        window.location.reload();
      });
    }

	
  },
  mounted() {
	
	if (this.isLoggedIn) {
	  this.showUserBanner = true;
	}
  },
};


</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
	  <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaText App</a>
	  <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
	  </button>
	</header>
  
	<div class="container-fluid">
	  <div class="row">
		<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		  <div class="position-sticky pt-3 sidebar-sticky">
			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
			  <span>General</span>
			</h6>
			<ul class="nav flex-column">
			  <li class="nav-item">
				<a class="nav-link active" href="#" @click.prevent="goHome">
				  <span data-feather="home"></span>
				  HOME
				</a>
			  </li>
			</ul>
		  </div>
		</nav>
		<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
		  <RouterView />
		</main>
	  </div>
	</div>
  
	<div v-if="isLoggedIn && showUserBanner" class="user-banner" @click="goToProfile">
	  <img :src="`data:image/jpg;base64,${user.photo}`" alt="Profile Picture" class="user-photo" />
	  <span class="user-name">{{ user.username }}</span>
	</div>
  </template>

<style>
.user-banner {
  position: fixed;
  bottom: 10px;
  left: 10px;
  display: flex;
  align-items: center;
  background-color: #c2c0c0b0;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  padding: 10px;
  gap: 10px;
  z-index: 1000;
}

.user-photo {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}
</style>

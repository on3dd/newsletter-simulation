<template>
    <div>
        <b-navbar variant="faded" type="light">
            <b-navbar-brand href="https://github.com/on3dd">
                <img src="https://avatars0.githubusercontent.com/u/45851782?s=460&v=4"
                     class="d-inline-block align-top brand-img"
                     alt="brand">
                on3dd
            </b-navbar-brand>
        </b-navbar>
    </div>
</template>

<script>
    export default {
        name: "Navbar"
    }
</script>

<style scoped>
    .brand-img {
        width: 30px;
        height: 30px;
        border-radius: 5px;
    }
</style>
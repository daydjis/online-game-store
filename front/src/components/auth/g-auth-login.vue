<template>
    <div class="auth-container ">
        <div class="auth-border">
            <form>
                <h2>Логин</h2>
                <div class="form-group">
                    <label for="exampleInputEmail1">Email address</label>
                    <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp"
                        placeholder="Enter email">
                    <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone
                        else.</small>
                </div>
                <div class="form-group">
                    <label for="exampleInputPassword1">Password</label>
                    <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
                </div>
                <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="exampleCheck1">
                    <label class="form-check-label" for="exampleCheck1">Check me out</label>
                </div>

                <router-link to="/auth/register">
                    <a class="nav-link active black" aria-current="page" href="#">Ещё не зарегистрировались?</a>
                </router-link>

                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.black {
    color: black
}

.auth-border {
    margin-top: 40px;
    box-shadow: 1px 2px 4px black;
    border-radius: 5px;
    padding: 20px;
}

.auth-container {
    display: flex;
    justify-content: center;
    align-items: center;
    max-width: 1200px;
    margin: 0 auto;
}
</style>
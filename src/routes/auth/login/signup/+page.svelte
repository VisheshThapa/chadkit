<script lang='ts'>
    import { pb, user } from "$lib/pocketbase";
    import { redirect } from "@sveltejs/kit";
    import {goto} from "$app/navigation"

import * as EmailValidator from 'email-validator';
    let email: string;
    let password: string;
    let passwordConfirm: string;
    async function signUp(){
        try{
        if(password == passwordConfirm && EmailValidator.validate(email)){
            console.log("if")
        console.log(email)
            console.log(password)

        const data = {
            email,
            password,
            passwordConfirm,
        };

        const record = await pb.collection('users').create(data);

            // (optional) send an email verification request
           await pb.collection('users').requestVerification(email);
            signIn()
            goto("/login/")

        }
        }
        catch(err) {
            console.log(err)
        }
    }
    async function signOut(){
        pb.authStore.clear();
    }
    async function signIn(){
        const authData = await pb.collection('users').authWithPassword(email, password);
    }
</script>

{#if $user}
  <h2 class="card-title">Welcome, {$user.email}</h2>
  <p class="text-center text-success">You are logged in</p>
  <p class="text-center text-success">Verified: {$user.verified}</p>
  <button class="btn btn-warning" on:click={() => signOut()}>Sign out</button>
{:else if $user == null}
    <p>Loading</p>
{:else}
    <div class="mt-8 flex items-center justify-center">

  <div class="card w-full max-w-sm shadow-2xl bg-base-100 ">
      <form class="card-body">
        <div class="form-control">
          <label class="label">
            <span class="label-text">Email</span>
          </label>
          <input type="email" placeholder="email" bind:value={email} class="input input-bordered" required />
        </div>
        <div class="form-control">
          <label class="label">
            <span class="label-text">Password</span>
          </label>
          <input type="password" placeholder="password" bind:value={password} class="input input-bordered" required />
                <label class="label">
            <span class="label-text"> Confirm Password</span>
          </label>
          <input type="password" placeholder="confirm password" bind:value={passwordConfirm} class="input input-bordered" required />
 <label class="label">
                            <a href="/login/" class="label-text-alt link link-hover">Login</a>
                        </label>
              
        <div class="form-control mt-6">
          <button class="btn btn-primary" on:click={signUp}>Sign Up</button>
        </div>
             </form>
        </div>
</div>
{/if}

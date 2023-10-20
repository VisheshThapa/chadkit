<script lang='ts'>
    import { pb, user } from "$lib/pocketbase";

    let email: string;
    let password: string;
    async function signIn(){
        const authData = await pb.collection('users').authWithPassword(email, password);
    }

    async function signInWithGoogle(){
        const authData = await pb.collection('users').authWithOAuth2({ provider: 'google' });
    }

    async function signOut(){
        pb.authStore.clear();
    }
    console.log(user)
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
            <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
          </label>
        <div class="form-control mt-6">
          <button class="btn btn-primary" on:click={signIn}>Login</button>
        </div>
                <div class="mb-3 mt-3">
                    <button on:click={signInWithGoogle} class="flex flex-wrap justify-center w-full border border-gray-300 hover:border-gray-500 px-2 py-1.5 rounded-md">
                        <img class="w-4 mr-4" src="https://lh3.googleusercontent.com/COxitqgJr1sJnIDe8-jiKhxDx1FrYbtRHKJ9z_hELisAlapwE9LUPh6fcXIfb5vwpbMl4xl9H9TRFPc5NOO8Sb3VSgIBrfRYvW6cUA">
                            Sign in with Google
                    </button>
                </div>
                        <label class="label">
                            <a href="/login/signup" class="label-text-alt link link-hover">Sign Up</a>
                        </label>
                    </div>
      </form>
        </div>
    </div>
{/if}


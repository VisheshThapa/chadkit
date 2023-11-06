<script lang="ts">
    import { goto } from "$app/navigation";
    function CheckoutBasic(){
        goto("https://buy.stripe.com/8wM00T7cNalNgne3cc");
    }
    function CheckoutPro(){
        goto("https://buy.stripe.com/aEUeVNdBb51tc6Y146")
    }
</script>
<div class  = "md:flex lg:gap-20 md:gap-10 grid justify-center content-center mt-5">
<div class="card card-bordered w-96 bg-base-100 shadow-xl mt-5">
  <div class="card-body">
    <h2 class="card-title">Chad</h2>
  </div>

    <div class="card-normal flex flex-row justify-center">
        <p class=" text-4xl  ">$120 </p>
        <p class="float-left">USD</p>
    </div>
        
    <div class="card-body ">
        <ul class="justify-center ">
            <li class= "flex "><img src="/checkmark.svg" alt = "checkmark" class = "max-h-4 self-center pr-2"/>Sveltekit</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Custom UI</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Pocketbase with Dockerfile</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Stripe</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Postmark</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Fly.io Guide</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>SEO</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>10% to Pocketbase's Opensource Community</li>
            <li class= "flex"><img src="/cancel.svg" class = "max-h-4 self-center pr-2"/>Full Setup Guide</li>
            <li class= "flex"><img src="/cancel.svg" class = "max-h-4 self-center pr-2"/>Latest Updates</li>
            <li class= "flex"><img src="/cancel.svg" class = "max-h-4 self-center pr-2"/>Discord Server</li>


        </ul>
        <div class="card-actions justify-center">
        <button class="btn btn-primary" on:click={CheckoutBasic}>Buy Now</button>
        </div>
    </div>
</div>

<!--Pro pricing -->
<div class="card card-bordered w-96 bg-base-100 shadow-xl mt-5">
  <div class="card-body">
    <h2 class="card-title">GigaChad</h2>
  </div>

    <div class="card-normal flex flex-row justify-center">
        <p class=" text-4xl  ">$173</p>
        <p class="float-left">USD</p>
    </div>
        
    <div class="card-body ">
        <ul class="justify-center ">
            <li class= "flex"><img src="/checkmark.svg" alt = "checkmark" class = "max-h-4 self-center pr-2"/>Sveltekit</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Custom UI</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Pocketbase with Dockerfile</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Stripe</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Postmark</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Fly.io Guide</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>SEO</li>
                <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>20% to Pocketbase's Opensource Community</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Full Setup Guide</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Latest Updates</li>
            <li class= "flex"><img src="/checkmark.svg" class = "max-h-4 self-center pr-2"/>Discord Server</li>

        </ul>
        <div class="card-actions justify-center">
        <button class="btn btn-primary" on:click={CheckoutPro}>Buy Now</button>
        </div>
    </div>
</div>
</div>

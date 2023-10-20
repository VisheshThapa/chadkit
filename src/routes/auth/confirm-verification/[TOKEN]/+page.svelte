<script lang="ts">
  import type { PageData } from "./$types";
    export let data : PageData;
    console.log(data.confirm);
</script>

{#if data.confirm}
    <h1>Thank you for confirming your email</h1>
{:else}
    <h1>Couldn't confirm your email, please try agian</h1>
{/if}


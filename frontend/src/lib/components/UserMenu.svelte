<script lang="ts">
import { steamIdFromToken } from '$lib/steam_login';
const { steamId }: { steamId?: string } = $props();

let steamIdState = $derived(steamId || (await steamIdFromToken()));
</script>

<div>
	<span>{steamIdState}</span>
</div>

<style lang="postcss">
</style>

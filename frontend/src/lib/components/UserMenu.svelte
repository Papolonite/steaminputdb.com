<script lang="ts">
import { page } from '$app/state';

// https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=81D64B225925B47B2EC27A9B14780D37&steamids=76561197997352479
</script>

<div>
	<span>{page.data.steamId}</span>
</div>

<style lang="postcss">
</style>

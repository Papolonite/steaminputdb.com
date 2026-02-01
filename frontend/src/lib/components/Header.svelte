<script lang="ts">
import { resolve } from '$app/paths';
import Themetoggle from '$lib/components/theme/toggle.svelte';
import { steamIdFromToken } from '$lib/steam_login';
import Icon from '@iconify/svelte';
import UserMenu from './UserMenu.svelte';

const steamId = $state(await steamIdFromToken());
</script>

<header>
	<a class="neutral" href={resolve('/')}><span>SteamInputDB.com</span></a>
	{#if !steamId}
		<a href={resolve('/login')}>
			<Icon icon="mdi:steam" width="1.2em" height="1.2em" />
			<span>Sign In</span>
		</a>
	{:else}
		<UserMenu steamId={steamId} />
	{/if}
	<Themetoggle />
</header>

<style lang="postcss">
header {
	padding: 1em;
	background: var(--card-color);
	display: grid;
	grid-template-columns: auto min-content min-content;
	align-items: center;
	box-shadow: 0 0px 4px var(--shadow-color);
	gap: 1em;
}

span {
	font-weight: bold;
}

.neutral {
	font-size: 1.4em;
	text-decoration: none;
	color: inherit;
	position: relative;
	width: min-content;
}

a {
	font-size: 1.4em;
	white-space: nowrap;
	display: grid;
	place-items: center;
	grid-auto-flow: column;
	gap: 0.25em;
	color: var(--text-color);
	&:hover,
	&:focus,
	&:focus-within {
		color: var(--color-primary);
	}
}
</style>

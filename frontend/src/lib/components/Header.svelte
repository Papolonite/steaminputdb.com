<script lang="ts">
import { beforeNavigate } from '$app/navigation';
import { resolve } from '$app/paths';
import { page } from '$app/state';
import Themetoggle from '$lib/components/theme/toggle.svelte';
import Icon from '@iconify/svelte';
import UserMenu from './UserMenu.svelte';

let loginRoute = $state(false);
beforeNavigate(({ from, to }) => {
	if (from?.route.id?.includes('/login') || to?.route.id?.includes('/login')) {
		loginRoute = true;
		return;
	}
	loginRoute = false;
});
</script>

<header>
	<a class="neutral" href={resolve('/')}><span>SteamInputDB.com</span></a>
	{#if !page.route.id?.startsWith('/login')}
		{#if !page.data.steamId}
			<a class={loginRoute ? 'login-view-transition' : ''} href={resolve('/login')}>
				<Icon icon="mdi:steam" width="1.2em" height="1.2em" />
				<span>Sign In</span>
			</a>
		{:else}
			<div class={loginRoute ? 'login-view-transition' : ''}>
				<UserMenu />
			</div>
		{/if}
	{/if}
	<Themetoggle />
</header>

<style lang="postcss">
header {
	padding: 1em;
	background: var(--card-color);
	display: flex;
	align-items: center;
	box-shadow: 0 0px 4px var(--shadow-color);
	gap: 1em;
	overflow: hidden;
	transition-property: all;
	& > :first-child {
		display: unset;
		margin-right: auto;
		overflow: hidden;
		text-overflow: ellipsis;
	}
}

.login-view-transition {
	view-transition-name: steamlogin;
}

::view-transition-old(steamlogin) {
	animation: var(--transition-duration) var(--ease-vanish) fade-out;
	animation-fill-mode: forwards;
}

::view-transition-new(steamlogin) {
	animation: var(--transition-duration) var(--ease-appear) slide-down-fade;
	animation-fill-mode: forwards;
}

@keyframes slide-down-fade {
	from {
		transform: translateY(-100%);
		opacity: 0;
	}
	to {
		transform: translateY(0);
		opacity: 1;
	}
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

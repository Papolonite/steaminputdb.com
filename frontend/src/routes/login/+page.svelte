<script lang="ts">
import { buildSteamLoginUrl } from '$lib/steam_login';
import Icon from '@iconify/svelte';
</script>

<main>
	<div>
		<p>
			Yadda yadda <br />
			Logging into SteamInputDB allows you to access additional features. To provide a better service, we
			will fetch public information about your account (steamid, profile name, avatar, and game list). This
			website is not affiliated with Valve Corporation or Steam. You can delete your account and all of the
			data at any time. We use cookies to keep your session active.
		</p>
		<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
		<a class="button" href={buildSteamLoginUrl()}
			>Sign in through <Icon icon="mdi:steam" width="1.4em" height="1.4em" /> Steam
		</a>
	</div>
</main>

<style lang="postcss">
main {
	padding: 2em;
	display: grid;
	place-items: center;
	width: 100%;
	height: 100%;
}

div {
	display: grid;
	place-items: center;
	gap: 2em;
}

a {
	border-radius: 100vh;
	font-size: 1.4em;
	& > :global(svg) {
		margin: 0 0.5ch;
	}
	width: fit-content;
	height: fit-content;
}
</style>

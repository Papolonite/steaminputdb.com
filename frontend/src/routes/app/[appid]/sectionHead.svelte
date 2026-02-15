<script lang="ts" module>
export { sectionHead };
</script>

<script lang="ts">
import type { components } from '$lib/api/openapi';
import { communityUrlBase, steamStoreUrlBase } from '$lib/steamapi/const';
import Icon from '@iconify/svelte';
import { cubicOut } from 'svelte/easing';
import { fade } from 'svelte/transition';
</script>

{#snippet sectionHead({
	appInfo,
	fallbackName
}: {
	appInfo?: components['schemas']['AppItem'];
	fallbackName?: string;
})}
	<section>
		<div>
			{#if appInfo?.assets?.community_icon}
				<picture transition:fade={{ duration: 196, easing: cubicOut }}>
					<enhanced:img
						src={`${communityUrlBase}${appInfo.app_id}/${appInfo.assets?.community_icon}.jpg`}
						alt="Icon"></enhanced:img>
				</picture>
			{:else}
				<Icon icon="mdi:link-variant" width="2.5em" height="2.5em" />
			{/if}
			{#if !appInfo && fallbackName}
				<i>(Non Steam Shortcut)</i>
			{/if}
			<h1>{appInfo?.name ?? fallbackName}</h1>
		</div>
		<div>
			<!-- 
            TODO: create buddy-app that interacts with steam via cef-remote-debug
            If you are reading this and think this works without - Nope CORS policy. and that's a good thing!
			<a href="#" class="button">
				<Icon icon="mdi:steam" width="1.4em" height="1.4em" />
				<span>Show Controller Config</span>
			</a> -->
			{#if appInfo?.store_url_path}
				<a
					href={steamStoreUrlBase + appInfo?.store_url_path}
					class="button"
					target="_blank"
					rel="external">
					<!-- <Icon icon="mdi:steam" width="1.4em" height="1.4em" /> -->
					<Icon icon="mdi:local-grocery-store" width="1.4em" height="1.4em" />
					<span>Storepage</span>
				</a>
			{/if}
		</div>
	</section>
{/snippet}

<style lang="postcss">
section {
	display: grid;
	max-width: 100%;
	gap: 1em;
	grid-template-columns: repeat(auto-fit, minmax(min(100%, 25ch), auto));

	padding: 0 1em;
	isolation: isolate;

	& > :first-child {
		margin: auto;
		display: grid;
		grid-template-columns: minmax(min-content, 2em) auto;
		align-items: center;
		width: 100%;
		height: fit-content;
		grid-column-gap: 1em;
		grid-row-gap: 0.25em;
		padding: 1em 0;

		& picture,
		& img {
			object-fit: cover;
			object-position: center;
			overflow: hidden;
			width: fit-content;
			box-shadow: 0 0.2em 0.7em 0em var(--shadow-color);
		}

		& > i {
			grid-row: 2 / span 1;
			grid-column: 1 / span 2;
		}

		& > :last-child {
			margin-right: auto;
			filter: drop-shadow(1px 1px 2px black);
		}
	}
	& > :last-child {
		display: grid;
		place-items: center;
		margin: auto;
		gap: 1em;
		width: 100%;
		grid-template-columns: repeat(auto-fit, minmax(19ch, auto));

		& > a {
			width: 100%;
			white-space: nowrap;
			display: grid;
			align-items: center;
			justify-content: center;
			gap: 0.5ch;
			font-weight: bold;
			background: linear-gradient(
				215deg,
				color-mix(in srgb, var(--card-color), transparent 35%) 0%,
				color-mix(in srgb, var(--card-color), transparent 60%) 70%
			);
			& > span {
				width: fit-content;
			}
		}
		& .button {
			&:hover,
			&:focus-visible {
				color: var(--text-color-dark);
				background-color: var(--color-primary);
			}
		}
		/* & .button:is(:first-child) {
			background-color: #1a9fff;
			&:hover,
			&:focus-visible {
				background-color: color-mix(in srgb, #1a9fff, var(--color-primary) 50%);
			}
		} */
	}
}
</style>

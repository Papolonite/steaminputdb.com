<script lang="ts" module>
export { sectionHead };
</script>

<script lang="ts">
import type { components } from '$lib/api/openapi';
import { tooltip } from '$lib/attachments/tooltip.svelte';
import { assetUrlBase, communityUrlBase } from '$lib/steamapi/const';
import Icon from '@iconify/svelte';
import { cubicOut } from 'svelte/easing';
import { fade } from 'svelte/transition';
</script>

{#snippet sectionHead({
	fileInfo,
	appInfo,
	isMobileBrowser
}: {
	fileInfo: components['schemas']['ConfigResponseItem'];
	appInfo?: components['schemas']['AppsSearchItem'];
	isMobileBrowser?: boolean;
})}
	<section>
		<div>
			{#if appInfo?.assets}
				{@const assets = appInfo?.assets}
				{@const assetChosen =
					assets.small_capsule ??
					assets.main_capsule ??
					assets.header ??
					assets.hero_capsule ??
					assets.library_hero ??
					'none.svg'}
				{#if assetChosen}
					<picture transition:fade={{ duration: 196, easing: cubicOut }}>
						<enhanced:img
							src={`${assetUrlBase}${assets.asset_url_format?.replace(
								'${FILENAME}',
								assetChosen
							)}`}
							alt="Capsule"
							height="100%"></enhanced:img>
					</picture>
				{:else}
					<div></div>
				{/if}
			{:else}
				<div></div>
			{/if}
			<div>
				<h1>{fileInfo.title}</h1>
				{#if appInfo}
					{#if appInfo.assets?.community_icon}
						<enhanced:img
							src={`${communityUrlBase}${appInfo.app_id}/${appInfo.assets.community_icon}.jpg`}
							alt="Icon"
							style="min-width: 1.2em; height: 1.2em; margin-right: 0.1em;"></enhanced:img>
					{:else}
						<Icon icon="mdi:steam" width="1.2em" />
					{/if}
				{:else}
					<Icon icon="mdi:link-variant" width="1.2em" />
				{/if}
				<h2>
					{appInfo?.name || fileInfo.app_id_string}
				</h2>
			</div>
		</div>
		<div>
			{#if true || !isMobileBrowser}
				{#snippet tooltipContent()}
					<div>
						<p style="white-space: nowrap; text-align: center;">Preview this config in Steam</p>
						<p>You must own the game</p>
						<code>steam://controllerconfig/{fileInfo.file_id}</code>
					</div>
				{/snippet}
				<a
					href={`steam://controllerconfig/${fileInfo.app_id_string}/${fileInfo.file_id}`}
					class="button blue"
					{@attach tooltip({
						snippet: tooltipContent,
						snippetInDefaultBackground: true,
						outDelay: 200,
						arrow: true,
						arrowFollowCursor: true
					})}>
					<Icon icon="mdi:steam" width="1.4em" height="1.4em" />
					<span>Preview | Apply</span>
				</a>
			{/if}

			{#if fileInfo.file_url}
				<a href={fileInfo.file_url} class="button" rel="external">
					<Icon icon="mdi:download" width="1.4em" height="1.4em" />
					<span>Download</span>
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

	& > :first-child {
		margin: auto;
		display: grid;
		place-items: center;
		width: 100%;
		height: fit-content;
		gap: 1em;
		padding: 1em 0;

		grid-template-columns: minmax(56px, min(420px, 33%)) auto;

		& > :first-child {
			min-height: 56px;
			height: 100%;
			width: 100%;
			background: linear-gradient(135deg, white -70%, transparent 120%);
			position: relative;

			& picture,
			& img {
				aspect-ratio: 21 / 8;
				object-fit: cover;
				object-position: center;
				width: 100%;
				overflow: hidden;
				z-index: -1;
			}

			z-index: -1;
			box-shadow: 0 0.2em 0.7em 0em var(--shadow-color);
		}

		& > :nth-child(2) {
			margin-right: auto;
			display: grid;
			height: fit-content;
			grid-template-columns: min-content auto;
			place-items: center;
			gap: 0.5ch;

			& > :first-child {
				grid-column: 1 / span 2;
				text-align: start;
				width: 100%;
				color: var(--text-color-dark);
				filter: drop-shadow(1px 1px 2px black);
			}
			& :global(> :nth-child(1n + 2)) {
				color: color-mix(in srgb, var(--color-primary), var(--text-color-dark) 60%);
				font-size: 1.8em;
				filter: drop-shadow(1px 1px 2px black);
			}
			& > :last-child {
				margin-right: auto;
			}
		}
	}
	& > :last-child {
		display: grid;
		place-items: center;
		margin: auto;
		gap: 1em;
		width: 100%;
		grid-template-columns: repeat(auto-fit, minmax(19ch, auto));

		& > .hov-over {
			width: 100%;
		}

		& > a {
			width: 100%;
			white-space: nowrap;
			display: grid;
			align-items: center;
			justify-content: center;
			gap: 0.5ch;
			font-weight: bold;
			border: 10px solid transparent;
			border: none !important;

			& > span {
				width: fit-content;
			}
		}
		& .button {
			background:
				linear-gradient(
					215deg,
					color-mix(in srgb, var(--card-color), transparent 75%) 0%,
					color-mix(in srgb, var(--card-color), transparent 90%) 70%
				),
				var(--bg-noise-transparent);
			&:hover,
			&:focus-visible {
				color: var(--text-color-dark);
				background-color: var(--color-primary);
			}
		}
		& .button:is(.blue) {
			background-color: #1a9fff;
			&:hover,
			&:focus-visible {
				background-color: color-mix(in srgb, #1a9fff, var(--color-primary) 50%);
			}
		}
	}
}

code {
	user-select: all;
	margin-top: 0.5em;
}
</style>

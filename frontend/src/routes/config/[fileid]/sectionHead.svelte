<script lang="ts" module>
export { sectionHead };
</script>

<script lang="ts">
import type { components } from '$lib/api/openapi';
import Icon from '@iconify/svelte';
import { cubicOut } from 'svelte/easing';
import { fade } from 'svelte/transition';
</script>

{#snippet sectionHead({
	fileInfo,
	appInfo
}: {
	fileInfo: components['schemas']['ConfigResponseItem'];
	appInfo?: components['schemas']['AppInfo'];
})}
	<section>
		<div>
			{#if appInfo?.header_image ?? appInfo?.capsule_imagev5 ?? appInfo?.capsule_image}
				{@const capsule = appInfo?.header_image ?? appInfo?.capsule_imagev5 ?? appInfo?.capsule_image}
				<picture transition:fade={{ duration: 196, easing: cubicOut }}>
					<enhanced:img src={capsule} alt="Thumbnail" height="100%"></enhanced:img>
				</picture>
			{:else}
				<div></div>
			{/if}
			<div>
				<h1>{fileInfo.title}</h1>
				{#if appInfo}
					<Icon icon="mdi:steam" width="1.2em" />
				{:else}
					<Icon icon="mdi:forbid" width="1.2em" />
				{/if}
				<h2>
					{appInfo?.name || fileInfo.app_id_string}
				</h2>
			</div>
		</div>
		<div>
			<a href={`steam://controllerconfig/${fileInfo.app_id_string}/${fileInfo.file_id}`} class="button">
				<Icon icon="mdi:steam" width="1.4em" height="1.4em" />
				<span>Preview | Apply</span>
			</a>
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
	isolation: isolate;

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
			& img {
				object-fit: cover;
				object-position: center;
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
		& .button:is(:first-child) {
			background-color: #1a9fff;
			&:hover,
			&:focus-visible {
				background-color: color-mix(in srgb, #1a9fff, var(--color-primary) 50%);
			}
		}
	}
}
</style>

<script lang="ts">
import { afterNavigate, beforeNavigate } from '$app/navigation';
import { resolve } from '$app/paths';
import { page } from '$app/state';
import { client } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import SC2Icon from '$lib/assets/SC2_Alt.svg?component';
import SC2 from '$lib/assets/SC2_Googley.svg.svelte';
import SearchForm from '$lib/components/search/searchform.svelte';
import Spinner from '$lib/components/Spinner.svelte';
import { log } from '$lib/log';
import { assetUrlBase } from '$lib/steamapi/const';
import Icon from '@iconify/svelte';
import { onMount } from 'svelte';
import { cubicIn, cubicInOut, cubicOut } from 'svelte/easing';
import { fade, slide } from 'svelte/transition';
import type { PageProps } from './$types';

let { data }: PageProps = $props();

let results = $derived.by(() => {
	if (data?.hasSearched) {
		return data.results;
	}
	return undefined;
});

let form = $state<HTMLFormElement>()!;
let loading = $state(false);
let hasSearched = $derived(data.hasSearched);
let searchError = $derived(data.searchError);

let eyes = $state<{
	left: HTMLElement;
	right: HTMLElement;
}>({
	left: undefined!,
	right: undefined!
})!;

let infoAppIdMap = $state<Record<number, components['schemas']['AppsSearchItem']>>({});

$effect(() => {
	results?.items?.forEach((cfg) => {
		if (!cfg.app_id) {
			return;
		}
		const idCopy = cfg.app_id;
		client
			.GET('/v1/steam/appinfo', {
				params: {
					query: {
						app_id: idCopy
					}
				}
			})
			.then((resp) => {
				if (resp.error) {
					log.error('Error fetching store info', 'appid', idCopy, 'error', resp.error);
					return;
				}
				if (!resp.data) {
					log.error('No data in response when fetching store info', 'appid', idCopy);
					return;
				}
				infoAppIdMap[idCopy] = resp.data as components['schemas']['AppsSearchItem'];
			})
			.catch((err) => {
				log.error('Error fetching store info', 'appid', idCopy, 'error', err);
			});
	});
});

let formValues = $state(
	Array.from(page.url.searchParams).reduce((acc, [key, value]) => {
		return Object.assign(acc, { [key]: value });
	}, {})
);

const findEyes = () => {
	const group = document.querySelector('#sc2>*>svg>g>:last-child');
	if (!group) {
		return;
	}
	eyes.left = group.children[0] as HTMLElement;
	eyes.right = group.children[1] as HTMLElement;
};
onMount(() => {
	findEyes();
});

beforeNavigate((event) => {
	console.log(event);
	if (event.type == 'form') {
		loading = true;
	}
});
afterNavigate(() => {
	loading = false;
});
</script>

<svelte:window
	onmousemove={(e) => {
		if (!eyes || !eyes.left || !eyes.right) {
			return;
		}
		Object.values(eyes).forEach((eye, idx) => {
			const rect = eye.getBoundingClientRect();
			const eyeX = rect.left + rect.width / 2;
			const eyeY = rect.top + rect.height / 2;
			const deltaX = e.clientX - eyeX;
			const deltaY = e.clientY - eyeY;
			const angle = Math.atan2(deltaY, deltaX);
			const distance = Math.min(16, Math.hypot(deltaX, deltaY) / 2);
			const translateX = Math.cos(angle) * distance * (searchError ? (idx === 0 ? 1 : -1) : 1);
			const translateY = Math.sin(angle) * distance * (searchError ? (idx === 0 ? -1 : 1) : 1);
			eye.style.transform = `translate(${translateX}px, ${translateY}px)`;
		});
	}} />

<main>
	<search>
		<SearchForm
			bind:form={form}
			disabled={loading}
			method="GET"
			bind:values={formValues}
			submitOnChange={true} />
		<div class="results">
			{#if loading}
				<div
					class="loading"
					in:fade|global={{ duration: 196, easing: cubicOut }}
					out:fade|global={{ duration: 196, easing: cubicIn }}>
					<Spinner size="16em" />
				</div>
			{/if}
			{#if !loading && (searchError || (results?.items?.length ?? 0) == 0)}
				<div
					id="sc2"
					class={searchError ? 'error' : ''}
					in:fade|global={{ duration: 196, easing: cubicOut }}
					out:fade|global={{ duration: 196, easing: cubicIn }}
					onintrostartcapture={findEyes}>
					{#if hasSearched && !searchError}
						<span>No results found</span>
					{/if}
					{#if searchError}
						<span>{searchError?.message ?? 'An unknown error occurred'}</span>
					{/if}
					<SC2
						height="100%"
						--eyes-color="black"
						--eyes-white-color="var(--text-color-dark)"
						--eyes-border-color="light-dark(var(--text-color-light), transparent)" />
				</div>
			{/if}
			{#if !searchError}
				<div class="list" transition:fade={{ duration: 196, easing: cubicInOut }}>
					{#each results?.items ?? [] as item (item.file_id)}
						<a
							class="plain"
							style={loading ? 'pointer-events: none; opacity: 0.4;' : ''}
							href={resolve(`/config/${item.file_id}`)}
							transition:slide={{ duration: 196, easing: cubicInOut }}>
							<div class="thumb">
								{#if infoAppIdMap?.[item.app_id || 0]}
									{@const assets = infoAppIdMap[item.app_id || 0]!.assets!}
									{@const srcChosen = assets?.asset_url_format
										? `${assetUrlBase}${assets?.asset_url_format?.replace(
												'${FILENAME}',
												assets.small_capsule ??
													assets.main_capsule ??
													assets.header ??
													assets.hero_capsule ??
													assets.library_hero ??
													assets.library_capsule ??
													assets.page_background ??
													'undefined'
											)}`
										: undefined}
									{#if srcChosen}
										<picture transition:fade={{ duration: 196, easing: cubicOut }}>
											<enhanced:img src={srcChosen} alt="Thumbnail" height="100%"
											></enhanced:img>
										</picture>
									{/if}
								{/if}
								<i>
									{#if item.app_id}
										<Icon icon="mdi:steam" width="1.2em" />
									{:else}
										<Icon icon="mdi:link-variant" width="1.2em" />
									{/if}
									{infoAppIdMap?.[item.app_id || 0]?.name ?? item.app_id_string}
								</i>
							</div>
							<div class="info">
								<strong class="title">{item.title}</strong>
								<span>
									{#if item.controller_type === 'controller_neptune'}
										<Icon icon="simple-icons:steamdeck" width="1.2em" />
									{:else if item.controller_type === 'controller_triton'}
										<SC2Icon width="1.2em" />
									{:else if item.controller_type === 'controller_steamcontroller_gordon'}
										<SC2Icon width="1.2em" />
									{:else if item.controller_type === 'controller_ps5'}
										<Icon icon="simple-icons:playstation5" width="1.2em" />
									{:else if item.controller_type === 'controller_ps4'}
										<Icon icon="iconoir:playstation-gamepad" width="1.2em" />
									{:else if item.controller_type === 'controller_xbox360'}
										<Icon icon="fluent:xbox-controller-24-regular" width="1.2em" />
									{:else if item.controller_type === 'controller_xboxone'}
										<Icon icon="fluent:xbox-controller-24-filled" width="1.2em" />
									{:else if item.controller_type === 'controller_switch_pro'}
										<Icon icon="mdi:controller" width="1.2em" />
									{:else if item.controller_type === 'controller_mobile_touch'}
										<Icon icon="mdi:cellphone" width="1.2em" />
									{:else if item.controller_type === 'controller_android'}
										<Icon icon="mdi:android" width="1.2em" />
									{:else}
										<Icon icon="mdi:gamepad" height="1.2em" />
									{/if}

									{item.controller_type_nice ||
										item.controller_type ||
										'Generic Controller'}
								</span>
							</div>
						</a>
					{/each}
				</div>
			{/if}
		</div>
	</search>
</main>

<style lang="postcss">
main {
	display: grid;
	place-items: center;
}

search {
	display: grid;
	place-self: center;
	grid-template-rows: min-content auto;
	gap: 1em;
	place-items: center;
	min-width: 50%;
	--max-width: 1440px;
	max-width: min(100%, var(--max-width));
	isolation: isolate;

	min-height: min(100%, 55dvh);

	& > :first-child {
		width: 100%;
	}
	margin-top: 1em;
}

.results {
	display: grid;
	place-items: center;
	width: 100%;
	grid-template-columns: auto;
	grid-template-rows: 1fr;
	padding: 1em 0;
	height: 100%;

	& > * {
		grid-row: 1;
		grid-column: 1;
	}
}

#sc2 {
	filter: drop-shadow(0px 0.25em 0.2em var(--shadow-color));
	width: 100%;
	display: grid;
	place-items: center;
	padding: 0.5em;
	gap: 1em;

	:global(svg) {
		max-height: 26em;
		max-width: min(100%, 26em);
		margin-bottom: auto;
	}

	@media (any-pointer: coarse) {
		:global(ellipse) {
			transition: transform calc(var(--transition-duration) * 2) var(--default-ease);
		}
	}

	& span {
		color: var(--highlight-color);
		font-size: 2em;
		font-weight: bold;
	}

	&.error {
		:global(svg) {
			transform: rotate(180deg);
		}

		:global(#qam),
		:global(#pad_r),
		:global(#y),
		:global(#x) {
			display: none;
		}

		:global(#logo_overlay) {
			fill: var(--background-color);
		}

		:global(#dpad) {
			transform: rotate(3deg);
			transform-origin: 25% 25%;
			translate: -0.1em 0.1em;
		}
	}
}

.results .list {
	width: 100%;
	display: grid;
	position: relative;
}

a {
	display: grid;
	--gap: 1em;
	padding: calc(0.5 * var(--gap)) calc(0.5 * var(--gap));
	column-gap: 1em;
	row-gap: 0;
	grid-template-columns: 33% 66%;
	color: var(--text-color);
	position: relative;
	isolation: isolate;
	&::before {
		content: '';
		height: 0.1em;
		width: 100%;
		bottom: calc(0.1em * -0.5);
		left: 0;
		position: absolute;
		background: color-mix(in srgb, currentColor, transparent 80%);
	}

	& > div {
		overflow: hidden;
		width: 100%;
		display: grid;
	}
}

.title {
	text-overflow: ellipsis;
	overflow: hidden;
	white-space: nowrap;
	font-size: 1.4em;
}

.thumb {
	aspect-ratio: 21 / 8;
	width: 100%;
	height: 100%;

	color: var(--text-color-dark);
	isolation: isolate;

	& picture,
	& img {
		object-fit: cover;
		object-position: center;
		width: 100%;
		height: 100%;
		overflow: hidden;
		z-index: -1;
	}

	& picture {
		isolation: isolate;
		position: relative;
		&::before {
			content: '';
			position: absolute;
			inset: 0px;
			mask: linear-gradient(
				135deg,
				rgb(255 255 255 / 0.9) 0%,
				rgb(255 255 255 / 0.7) 25%,
				transparent 50%
			);
			mask-type: alpha;
			backdrop-filter: saturate(0.1) brightness(0.3);
			z-index: 1;
		}
	}

	display: grid;
	background: linear-gradient(rgb(90 90 90 / 1), rgb(150 150 150 / 0.9));
	& > * {
		grid-row: 1;
		grid-column: 1;
	}

	& i {
		font-size: 1.1em;
		padding: 0.5em;
		filter: drop-shadow(1px 1px 1px black) drop-shadow(0px 0px 4px rgb(0 0 0 / 0.25))
			drop-shadow(0px 0px 2px black);
		display: grid;
		grid-auto-flow: column;
		align-items: center;
		justify-content: start;
		height: fit-content;
		gap: 0.25em;
		font-weight: bold;
	}
}

.info {
	display: grid;
	grid-template-rows: min-content min-content auto;
	& span {
		display: grid;
		align-items: center;
		grid-auto-flow: column;
		width: fit-content;
		gap: 0.5em;
		font-size: 1.2em;
	}
	strong {
		padding-bottom: 0.5em;
	}

	& span {
		color: var(--highlight-color);
		font-weight: bold;
	}
}

.loading {
	align-self: baseline;
	z-index: 1;
}
</style>

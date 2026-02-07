<script lang="ts">
import { client, type ResponseType } from '$lib/api/client';
import { Debounced } from '$lib/debounce.svelte';
import { log } from '$lib/log';
import { cubicIn, cubicOut } from 'svelte/easing';
import { slide } from 'svelte/transition';
import Searchbar from '../search/searchbar.svelte';
import SearchResults from './SearchResults.svelte';

let form = $state<HTMLFormElement>()!;
let debounced = new Debounced<string>(223).eager();
let focusState = new Debounced(96, false).eager();

let previewResults = $state<ResponseType<'POST', '/v1/search/'>['data']>({
	games: [],
	configs: []
});

let shouldShowPreview = $derived.by(() => {
	if (!focusState.value) {
		return false;
	}
	if ((previewResults?.games?.length || 0) === 0 && (previewResults?.configs?.length || 0) === 0) {
		return false;
	}
	return true;
});

$effect(() => {
	fetchLivePreview(debounced.value);
});

const fetchLivePreview = (search_term = '') => {
	debounced.lock();
	client
		.POST('/v1/search/', {
			body: {
				search_term,
				limit_games: 3,
				limit_configs: 5,
				limit: 10
			}
		})
		.then((res) => {
			if (res.error) {
				log.error('Error fetching search preview', 'err', res.error);
				return;
			}
			if (res.data) {
				previewResults = res.data;
			}
		})
		.catch((err) => {
			log.error('Error fetching search preview', 'err', err);
		})
		.finally(() => {
			debounced.unlock(true);
		});
};
if ((previewResults?.games?.length || 0) === 0 || (previewResults?.configs?.length || 0) === 0) {
	void fetchLivePreview();
}
</script>

<search onfocusin={() => (focusState.input = true)} onfocusout={() => (focusState.input = false)}>
	<!-- TODO: change endpoint -->
	<form bind:this={form} method="POST" action="/?/search">
		<Searchbar
			--box-shadow="inset 0 0.4em 0.4em 0 var(--shadow-color-dark)"
			bind:value={debounced.input} />
	</form>
	{#if shouldShowPreview}
		<dialog
			open
			in:slide={{ axis: 'y', duration: 196, easing: cubicOut }}
			out:slide={{ axis: 'y', duration: 196, easing: cubicIn }}>
			<SearchResults results={previewResults} />
		</dialog>
	{/if}
</search>

<style lang="postcss">
search {
	position: relative;
	isolation: isolate;
}

form {
	place-self: center;
	display: grid;
	grid-template-columns: minmax(16ch, 48ch);
	z-index: 1;
}

dialog[open] {
	--oversize: 50%;
	--corner-radius: 1em;
	position: absolute;
	isolation: isolate;
	top: 100%;
	left: 50%;
	translate: -50% 0;
	width: calc(100% + var(--oversize));
	max-width: 100dvw;

	padding-top: 0.5em;

	border: none;
	border-radius: 0 0 var(--corner-radius) var(--corner-radius);
	margin-top: 0.5em;
	background: var(--card-background-noise);
	background-color: var(--card-color);
	box-shadow: 0 0.5em 1em -0.5em var(--shadow-color);

	&::before {
		@supports (corner-shape: scoop) {
			content: '';
		}
		position: absolute;
		top: 0.5em;
		left: calc(var(--corner-radius) * -1);
		corner-shape: scoop;
		border-radius: 0 0 0 var(--corner-radius);
		width: var(--corner-radius);
		height: var(--corner-radius);
		background: var(--card-background-noise);
		background-color: var(--card-color);
		z-index: -1;
	}
	&::after {
		@supports (corner-shape: scoop) {
			content: '';
		}
		position: absolute;
		top: 0.5em;
		right: calc(var(--corner-radius) * -1);
		corner-shape: scoop;
		border-radius: 0 0 var(--corner-radius) 0;
		width: var(--corner-radius);
		height: var(--corner-radius);
		background: var(--card-background-noise);
		background-color: var(--card-color);

		z-index: -1;
	}
}
</style>

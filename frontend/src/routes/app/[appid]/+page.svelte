<script lang="ts">
import { page } from '$app/state';
import type { components } from '$lib/api/openapi';
import { assetUrlBase, storePageBackgroundBase } from '$lib/steamapi/const';
import { sectionHead } from './sectionHead.svelte';

const appInfo: components['schemas']['AppItem'] = $derived(page.data.appInfo);

const pageBGURL = $derived.by(() => {
	if (!appInfo?.assets) {
		return;
	}
	if (appInfo.assets.page_background) {
		return `${assetUrlBase}${appInfo.assets.asset_url_format?.replace(
			'${FILENAME}',
			appInfo.assets.page_background
		)}`;
	}
	if (appInfo.assets.raw_page_background) {
		return `${storePageBackgroundBase}${appInfo.assets.asset_url_format?.replace(
			'${FILENAME}',
			appInfo.assets.raw_page_background
		)}`;
	}
	if (appInfo.assets.page_background_path) {
		return `${storePageBackgroundBase}${appInfo.assets.page_background_path}`;
	}
});
</script>

<main style={pageBGURL ? `--bg: url('${pageBGURL}')` : ''}>
	<div>
		{@render sectionHead({ appInfo })}
		<!-- TODO: show controller support and whatnot -->
	</div>
</main>

<style lang="postcss">
main {
	position: relative;
	isolation: isolate;
	display: grid;
	padding: 1em 0;

	place-items: center;
	grid-template-rows: min-content;
	grid-template-columns: minmax(min(100%, auto), 50%);
	width: 100%;

	&::before {
		content: '';
		position: absolute;
		inset: 0;
		background:
			linear-gradient(
				0deg,
				var(--background-color) 25%,
				color-mix(in srgb, var(--background-color), transparent 100%) 80%
			),
			var(--bg, transparent) top/cover no-repeat;
		z-index: -2;
	}
}
div {
	display: grid;
	place-self: center;
	gap: 1em;
	place-items: center;
	min-width: 50%;
	--max-width: 1440px;
	max-width: min(100%, var(--max-width));
	isolation: isolate;
	/* container: main / inline-size;*/
	:global(> :first-child) {
		z-index: -1;
		width: 100%;
	}
}
</style>

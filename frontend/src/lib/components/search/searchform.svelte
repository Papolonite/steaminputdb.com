<script lang="ts">
import { enhance } from '$app/forms';
import SC2 from '$lib/assets/SC2_Alt.svg?component';
import Searchbar from '$lib/components/search/searchbar.svelte';
import Icon from '@iconify/svelte';
import { tick } from 'svelte';
import { cubicInOut } from 'svelte/easing';
import type { HTMLFormAttributes } from 'svelte/elements';
import { slide } from 'svelte/transition';

let {
	showControllerFilter = true,
	showFeatureFilter = true,
	form = $bindable(),
	searchtext = $bindable(),
	disabled = false,
	method = 'GET',
	values = $bindable({}),
	submitOnChange = false,
	enhanceParams,
	...rest
}: {
	showControllerFilter?: boolean;
	showFeatureFilter?: boolean;
	form?: HTMLFormElement;
	searchtext?: string;
	disabled?: boolean;
	method?: string;
	values?: Record<string, unknown>;
	submitOnChange?: boolean;
	enhanceParams?: Parameters<typeof enhance>[1];
} & HTMLFormAttributes = $props();

const changeSubmitHandler = () => {
	if (submitOnChange) {
		form!.requestSubmit();
	}
};
</script>

{#if method === 'POST' || method === 'post'}
	<form bind:this={form} class="card glass" {...rest} use:enhance={enhanceParams}>
		{@render formcontents()}
	</form>
{:else}
	<form bind:this={form} class="card glass" {...rest}>
		{@render formcontents()}
	</form>
{/if}

{#snippet formcontents()}
	<div>
		<Searchbar
			name="searchtext"
			placeholder="SteamInput configuration..."
			disabled={disabled}
			bind:value={values.searchtext}
			inlineButton={false} />
		<button type="submit" disabled={disabled}>Search</button>
		<label for="sort-by">
			<span>Sort by:</span>
			<select
				id="sort-by"
				name="sort-by"
				disabled={disabled}
				bind:value={values['sort-by']}
				onchange={changeSubmitHandler}>
				<option value="vote">Rank</option>
				<option value="publication">Date</option>
				<option value="trend">Trend (30 days)</option>
				<option value="votes_asc">Votes (ascending)</option>
				<option value="votes_up">Votes (upvotes)</option>
				<option value="text_search">Relevance</option>
				<option value="playtime_trend">Playtime trend (30 days)</option>
				<option value="total_playtime">Total playtime</option>
				<option value="avg_playtime_trend">Average playtime trend</option>
				<option value="lifetime_avg_playtime">Average playtime since upload</option>
				<option value="playtime_sessions_trend">Sessions trend (30 days)</option>
				<option value="lifetime_playtime_sessions">Lifetime sessions</option>
			</select>
			<Icon icon="mdi:chevron-down" />
		</label>
	</div>
	{#if showControllerFilter}
		<fieldset
			id="controller-type"
			transition:slide={{ duration: 196, easing: cubicInOut }}
			onclickcapture={(e) => {
				const target = e.target;
				if (!(target instanceof HTMLInputElement)) {
					return;
				}
				if (target.type !== 'radio') {
					return;
				}
				if (target.name !== 'controller_type') {
					return;
				}

				if (values['controller_type'] == target.value) {
					values['controller_type'] = undefined;
					tick().then(() => {
						changeSubmitHandler();
					});
				}
			}}
			disabled={disabled}>
			<legend>Controller Type</legend>

			<label for="controller_neptune">
				<input
					type="radio"
					id="controller_neptune"
					name="controller_type"
					value="controller_neptune"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="simple-icons:steamdeck" width="1.2em" />
				<span> Steam Deck </span>
			</label>
			<label for="controller_triton">
				<input
					type="radio"
					id="controller_triton"
					name="controller_type"
					value="controller_triton"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<SC2 width="1.2em" />
				<span> Steam Controller </span>
			</label>
			<label for="controller_steamcontroller_gordon">
				<input
					type="radio"
					id="controller_steamcontroller_gordon"
					name="controller_type"
					value="controller_steamcontroller_gordon"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<SC2 width="1.2em" />
				<span> Steam Controller (2015) </span>
			</label>
			<label for="controller_ps5">
				<input
					type="radio"
					id="controller_ps5"
					name="controller_type"
					value="controller_ps5"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="simple-icons:playstation5" width="1.2em" />
				<span> DualSense </span>
			</label>
			<label for="controller_ps4">
				<input
					type="radio"
					id="controller_ps4"
					name="controller_type"
					value="controller_ps4"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="iconoir:playstation-gamepad" width="1.2em" />
				<span> DualShock 4 </span>
			</label>
			<label for="controller_xbox360">
				<input
					type="radio"
					id="controller_xbox360"
					name="controller_type"
					value="controller_xbox360"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="fluent:xbox-controller-24-regular" width="1.2em" />
				<span> Xbox 360 </span>
			</label>
			<label for="controller_xboxone">
				<input
					type="radio"
					id="controller_xboxone"
					name="controller_type"
					value="controller_xboxone"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="fluent:xbox-controller-24-filled" width="1.2em" />
				<span> Xbox One </span>
			</label>
			<label for="controller_switch_pro">
				<input
					type="radio"
					id="controller_switch_pro"
					name="controller_type"
					value="controller_switch_pro"
					bind:group={values['controller_type'] as string}
					onchange={changeSubmitHandler} />
				<Icon icon="mdi:controller" width="1.2em" />
				<span> Switch Pro </span>
			</label>
		</fieldset>
	{/if}
	{#if showFeatureFilter}
		<fieldset id="features" transition:slide={{ duration: 196, easing: cubicInOut }} disabled={disabled}>
			<legend>Must have</legend>
			<label for="feature_gamepad">
				<input
					type="checkbox"
					id="feature_gamepad"
					name="feature_gamepad"
					bind:checked={values['feature_gamepad'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="mdi:controller" width="1.2em" />
				<span>Gamepad Inputs</span>
			</label>
			<label for="feature_keyboard">
				<!-- actually typo in valves data: feature_keboard -->
				<input
					type="checkbox"
					id="feature_keyboard"
					name="feature_keboard"
					bind:checked={values['feature_keboard'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="mdi:keyboard" width="1.2em" />
				<span>Keyboard Inputs</span>
			</label>
			<label for="feature_mouse">
				<input
					type="checkbox"
					id="feature_mouse"
					name="feature_mouse"
					bind:checked={values['feature_mouse'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="mdi:mouse" width="1.2em" />
				<span>Mouse Inputs</span>
			</label>
			<label for="feature_gyro">
				<input
					type="checkbox"
					id="feature_gyro"
					name="feature_gyro"
					bind:checked={values['feature_gyro'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="game-icons:gyroscope" width="1.2em" />
				<span>Gyro Inputs</span>
			</label>
			<label for="feature_touchmenu">
				<input
					type="checkbox"
					id="feature_touchmenu"
					name="feature_touchmenu"
					bind:checked={values['feature_touchmenu'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="mdi:gesture-touch" width="1.2em" />
				<span>Touch Menus</span>
			</label>
			<label for="feature_radialmenu">
				<input
					type="checkbox"
					id="feature_radialmenu"
					name="feature_radialmenu"
					bind:checked={values['feature_radialmenu'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="material-symbols:joystick" width="1.2em" />
				<span>Radial Menus</span>
			</label>
			<label for="feature_modeshift">
				<input
					type="checkbox"
					id="feature_modeshift"
					name="feature_modeshift"
					bind:checked={values['feature_modeshift'] as boolean}
					onchange={changeSubmitHandler} />
				<Icon icon="material-symbols:layers-rounded" width="1.2em" />
				<span>Mode Shifts</span>
			</label>
		</fieldset>
	{/if}
{/snippet}

<style lang="postcss">
form {
	display: flex;
	flex-flow: row wrap;
	width: 100%;
	gap: 1em;

	max-width: 100dvw;

	& > :first-child {
		width: 100%;
		flex-grow: 1;
		gap: 1em;
		width: 100%;
		display: flex;
		flex-flow: row wrap-reverse;
		position: relative;
		margin-bottom: 1em;
		:global(> :first-child) {
			flex-grow: 1;
			max-width: max(52ch, 25dvw);
		}
	}

	& > :nth-child(2) {
		flex: 1;
	}
	& > :nth-child(3) {
		flex: 1;
	}

	label {
		display: flex;
		gap: 0.5em;
		align-items: center;
	}

	label[for='sort-by'] {
		margin-left: auto;
		display: grid;
		grid-template-columns: auto auto;
		gap: 0.5em;
		align-items: center;
		font-size: 1.2em;
		position: relative;
		isolation: isolate;
		border: 1px solid color-mix(in srgb, var(--text-color), transparent 90%);
		padding: 0.5em 1em;
		box-shadow: 0 1px 4px 0 rgb(0 0 0 / 0.25);
		border-radius: 0.5em;
		transition: all var(--transition-duration) var(--default-ease);

		&[disabled] {
			opacity: 0.5;
		}

		&:has([disabled]) {
			opacity: 0.5;
		}

		&:hover,
		&:focus-within {
			outline: 0.1em solid var(--color-primary);
			box-shadow: 0 0 1.3em -0.4em var(--color-primary);
		}

		& > :first-child {
			white-space: nowrap;
		}
		:global(> :last-child) {
			content: '';
			color: var(--text-color);
			position: absolute;
			z-index: 1;
			height: 100%;
			width: 1.4em;
			top: 50%;
			translate: 0 -50%;
			right: 0.5em;
			background-size: contain;
			pointer-events: none;
		}
	}
}

select {
	font-style: inherit;
	background: transparent;
	border: 1px solid transparent;
	outline: none;
	color: var(--text-color);
	cursor: pointer;
	appearance: none;
	padding-right: 2em;
	position: relative;
	width: 100%;

	& option {
		color: var(--text-color);
		background: var(--card-color);
	}
}

fieldset {
	border-radius: 0.5em;
	padding: 1em;
	background: var(--card-background-noise);
	border: 1px solid color-mix(in srgb, var(--text-color), transparent 90%);
	position: relative;
	box-shadow: inset 0.1em 0.2em 0.5em 0 light-dark(#0f0f0f27, #0e0e0e7e);

	&[disabled] {
		opacity: 0.5;
	}

	& legend {
		font-size: 1.1em;
		border-radius: 0.5em;
		background: var(--card-background-noise);
		padding: 0.25em 0.5em;
		isolation: isolate;
		position: relative;
		&::before {
			content: '';
			position: absolute;
			inset: 0;
			z-index: -1;
			border-radius: inherit;
			border: 1px solid transparent;
			background: linear-gradient(0deg, transparent 0%, transparent 40%, var(--text-color) 100%)
				border-box;
			mask:
				linear-gradient(black, black) border-box,
				linear-gradient(black, black) padding-box;
			mask-composite: subtract;
			opacity: 0.2;
		}
	}
}

#controller-type {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(25ch, 1fr));
	gap: 1em;
	position: relative;
	inset: 0;
	& input {
		min-width: 1.4em;
		min-height: 1.4em;
	}

	& label {
		display: grid;
		grid-template-columns: min-content min-content auto;
		align-items: center;
		gap: 0.5em;
	}
}
#features {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(22ch, 1fr));
	gap: 1em;
	position: relative;
	inset: 0;
	& input {
		min-width: 1.4em;
		min-height: 1.4em;
	}

	& label {
		display: grid;
		grid-template-columns: min-content min-content auto;
		align-items: center;
		gap: 0.5em;
	}
}

button {
	color: var(--text-color-dark);
	font-weight: bold;
	background:
		linear-gradient(
			215deg,
			color-mix(in srgb, var(--card-color), transparent 75%) 0%,
			color-mix(in srgb, var(--card-color), transparent 90%) 70%
		),
		var(--bg-noise-transparent);
	background-color: color-mix(in srgb, var(--color-primary), transparent 20%);

	&[disabled] {
		opacity: 0.5;
	}

	&:hover,
	&:focus-visible {
		color: var(--text-color-dark) !important;
		background-color: color-mix(in srgb, var(--color-primary), rgb(128 128 255 / 0.8) 50%);
	}
}
</style>

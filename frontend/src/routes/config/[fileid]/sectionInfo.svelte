<script lang="ts" module>
export { sectionInfo };
</script>

<script lang="ts">
import { resolve } from '$app/paths';
import type { components } from '$lib/api/openapi';
import SC2 from '$lib/assets/SC2_Alt.svg?component';
import { selectAllHandler } from '$lib/attachments/selectAllHandler.svelte';
import Icon from '@iconify/svelte';
import { format, formatDistanceToNow, formatDuration, intervalToDuration } from 'date-fns';
</script>

{#snippet sectionInfo({
	fileInfo,
	appInfo,
	creatorInfo
}: {
	fileInfo: components['schemas']['ConfigDetailResponse'];
	appInfo?: components['schemas']['AppInfo'];

	creatorInfo?: components['schemas']['PlayerInfo'];
})}
	<section id="info">
		<dl>
			<dt>Controller</dt>
			<dd>
				{#if fileInfo.controller_type === 'controller_neptune'}
					<Icon icon="simple-icons:steamdeck" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_triton'}
					<SC2 width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_steamcontroller_gordon'}
					<SC2 width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_ps5'}
					<Icon icon="simple-icons:playstation5" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_ps4'}
					<Icon icon="iconoir:playstation-gamepad" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_xbox360'}
					<Icon icon="fluent:xbox-controller-24-regular" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_xboxone'}
					<Icon icon="fluent:xbox-controller-24-filled" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_switch_pro'}
					<Icon icon="mdi:controller" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_mobile_touch'}
					<Icon icon="mdi:cellphone" width="1.2em" />
				{:else if fileInfo.controller_type === 'controller_android'}
					<Icon icon="mdi:android" width="1.2em" />
				{:else}
					<Icon icon="mdi:gamepad" height="1.2em" />
				{/if}

				{fileInfo.controller_type_nice ?? fileInfo.controller_type ?? 'Generic Controller'}
			</dd>
			{#if creatorInfo}
				<dt>Author</dt>
				<dd>
					{creatorInfo.personaname}
				</dd>
			{/if}
			<dt>Author ID</dt>
			<dd>
				{fileInfo.creator_id}
			</dd>
			<dt>Created</dt>
			<dd>
				{format(new Date(fileInfo.time_created), 'PPpp')}
				<i>({formatDistanceToNow(new Date(fileInfo.time_created))} ago)</i>
			</dd>
			{#if fileInfo.file_url}
				<dt>File URL</dt>
				<dd>
					<code
						{@attach selectAllHandler(
							`outline: 1px solid transparent;
                                        background: rgb(128 128 128 / 0.10);`
						)}>{fileInfo.file_url}</code>
				</dd>
			{/if}
			{#if fileInfo.file_name}
				<dt>File Name</dt>
				<dd>
					<code
						{@attach selectAllHandler(
							`outline: 1px solid transparent;
                                        background: rgb(128 128 128 / 0.10);`
						)}>{fileInfo.file_name}</code>
				</dd>
			{/if}
			{#if fileInfo.file_size}
				<dt>File Size</dt>
				<dd>{(fileInfo.file_size / 1000).toFixed(0)} kB</dd>
			{/if}

			{#if fileInfo.playtime_seconds}
				{@const duration = intervalToDuration({
					start: new Date(0),
					end: new Date(fileInfo.playtime_seconds * 1000)
				})}
				<dt>Playtime</dt>
				<dd>
					{formatDuration(duration, {
						zero: false,
						format: ['years', 'months', 'days', 'hours']
					})}
					{#if (duration.hours || 0) > 0}
						and {formatDuration(duration, { format: ['minutes'] })}
					{/if}
					<i>(last two weeks)</i>
				</dd>
			{/if}
			{#if fileInfo.playtime_sessions}
				<dt>Sessions</dt>
				<dd>
					{fileInfo.playtime_sessions.toLocaleString()}
					<i>(last two weeks)</i>
				</dd>
			{/if}
			{#if fileInfo.lifetime_playtime_seconds}
				{@const duration = intervalToDuration({
					start: new Date(0),
					end: new Date(fileInfo.lifetime_playtime_seconds * 1000)
				})}
				<dt>Total</dt>
				<dd>
					<dl>
						<dt>Playtime</dt>
						<dd>
							{formatDuration(duration, {
								zero: false,
								format: ['years', 'months', 'days', 'hours']
							})}
							{#if (duration.hours || 0) > 0}
								and {formatDuration(duration, { format: ['minutes'] })}
							{/if}
						</dd>
						{#if fileInfo.lifetime_playtime_sessions}
							<dt>Sessions</dt>
							<dd>
								{fileInfo.lifetime_playtime_sessions.toLocaleString()}
							</dd>
						{/if}
					</dl>
				</dd>
			{/if}

			<dt>{appInfo ? 'Game' : 'Shortcut'}</dt>
			<dd>
				{#if appInfo}
					<a href={resolve(`/app/${appInfo.steam_appid}`)}>
						{appInfo.name} <i>({appInfo.steam_appid})</i></a>
				{:else}
					{fileInfo.app_id_string}
				{/if}
			</dd>
		</dl>
		<aside>
			<section>
				<div>Rating</div>
				<div>Playtime</div>
			</section>
			<p>
				<span>Description</span>
				{fileInfo.description?.replace(/\s\s/g, '\n')}
			</p>
		</aside>
	</section>
{/snippet}

<style lang="postcss">
#info {
	display: flex;
	flex-flow: row wrap-reverse;

	--gap: 1em;
	--info-min-width: 58ch;
	width: 100%;

	gap: 1em;
	padding: 0 1em;

	& > :first-child {
		border-radius: var(--border-radius);
		background: var(--card-glass);
		box-shadow: var(--card-shadow);
		padding: 1em;
		overflow-x: hidden;

		& > dt {
			font-size: 1.2em;
		}

		flex-basis: calc(66% - var(--gap));
		min-width: min(100%, var(--info-min-width));
	}
	& > :last-child {
		display: grid;
		height: fit-content;
		align-self: flex-end;
		width: 100%;
		flex-shrink: 1;
		gap: 1em;

		--width: 20ch;

		flex-basis: calc(34%);
		min-width: min(100%, var(--width));

		& > :first-child {
			display: grid;
			place-items: center;
			margin: auto;
			gap: 1em;
			width: 100%;
			grid-template-columns: repeat(auto-fit, minmax(calc(var(--width) -1em), auto));
			& > * {
				width: 100%;
				padding: 1em;
				position: relative;
				isolation: isolate;
				border-radius: var(--border-radius);
				box-shadow: var(--card-shadow);
				&::before {
					content: '';
					position: absolute;
					inset: 0;
					background: var(--card-glass);
					opacity: 0.5;
				}
			}
		}
	}

	@media (max-width: 200ch) {
		/* @container main (width < 800px) { */
		& > :first-child {
			flex-basis: 0;
			flex: 1;
			flex-grow: 1;
		}
		& > :last-child {
			flex-basis: 100%;
		}
	}
}

dl {
	display: grid;
	place-items: center;
	grid-template-columns: min-content auto;
	grid-column-gap: 1em;
	grid-row-gap: 0.5em;

	& > :nth-child(2) {
		font-weight: bold;
		:global(svg) {
			translate: 0 0.2em;
		}
	}
	& > :nth-child(2n-1) {
		justify-self: end;
		white-space: nowrap;
	}
	& > :nth-child(2n) {
		justify-self: start;
		white-space: no-wrap;
	}
}

dd {
	max-width: 100%;
	overflow: auto;
	overflow-x: hidden;
	& a {
		font-weight: bold;
	}
}

dt {
	font-weight: bold;
}

dd > dl {
	grid-column-gap: 0.5em;
	& > dt {
		color: var(--text-muted);
	}
}

dl {
	& i {
		color: var(--text-muted);
		white-space: nowrap;
	}
}

aside {
	& > p {
		overflow: auto;
		overflow-x: hidden;

		border-radius: var(--border-radius);
		background: var(--card-glass);
		box-shadow: var(--card-shadow);
		padding: 1em;
	}
}
</style>

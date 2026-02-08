<script lang="ts" module>
export { sectionInfo };
</script>

<script lang="ts">
import { resolve } from '$app/paths';

import type { components } from '$lib/api/openapi';
import { selectAllHandler } from '$lib/attachments/selectAllHandler.svelte';
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
			<dd>{fileInfo.controller_type_nice ?? fileInfo.controller_type ?? 'Generic Controller'}</dd>
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
						delimiter: ' ',
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
								delimiter: ' ',
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
			<section></section>
			<p></p>
		</aside>
	</section>
{/snippet}

<style lang="postcss">
section {
	display: flex;
	flex-flow: row wrap;
	gap: 1em;
	height: fit-content;
	place-content: center;
	max-width: 100%;

	& > :first-child {
		border-radius: var(--border-radius);
		background: var(--card-glass);
		box-shadow: var(--card-shadow);
		padding: 1em;
		max-width: 100%;
		overflow-x: hidden;
		& > dt {
			font-size: 1.2em;
		}
	}

	& > :last-child {
		display: grid;
		& > :first-child {
			display: grid;
		}
	}
}

dl {
	display: grid;
	place-items: center;
	grid-template-columns: min-content auto;
	grid-column-gap: 1em;
	grid-row-gap: 0.5em;

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
</style>

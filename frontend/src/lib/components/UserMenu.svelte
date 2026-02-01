<script lang="ts">
// eslint-disable-next-line prettier/prettier
let profileId: string | undefined = $state(
    (await cookieStore?.get('steamid'))
        ?.value
    );
</script>

<div>
	<span>{profileId}</span>
</div>

<style lang="postcss">
</style>

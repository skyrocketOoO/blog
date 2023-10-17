<script>
	import { onMount } from 'svelte';
	export let stories = [];
    onMount(async () => {
		const res = await axios.get("http://localhost:8080/story/");
		stories = await res.data["context"];
	});
</script>

<main>
    <div class="midcolumn">
        {#each stories as story}
        <div class="card">
            <h2>{story["Title"]}</h2>
            <h3>{story["Context"]}</h3>
            <p>{story["Label"]}</p>
            <h5>{story["Author"]}</h5>
        </div>
        {:else}
        <div class="card">
            <h2>文章标题</h2>
            <h5>2019 年 4 月 17日</h5>
            <div class="fakeimg" style="height:200px;">图片</div>
            <p>一些文本...</p>
            <p>菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！</p>
        </div>
        {/each}
    </div>
</main>

<style>
	.midcolumn {   
        float: left;
        width: 72%;
        overflow: scroll;
        background-color: aliceblue;
        padding: 14px;
        height: auto;
    }
</style>
<script>
	export let name = 'tasha';
	import axios from 'axios';
	
	function toggleName(){
		if (name === 'tasha'){
			name = 'zelin'
		}else{
			name = 'tasha'
		}
	}

	import { onMount } from 'svelte';
	let quotes = [];
    let error = null;
    onMount(async () => {
        try {
            const res = await axios.get('', {
                headers: {
                    'x-rapidapi-host': 'famous-quotes4.p.rapidapi.com',
                    'x-rapidapi-key': API_KEY,
                },
                params: { category: 'all', count: '10' },
            });
            quotes = res.data;
        } catch (e) {
            error = e;
        }
    });

	let src="images/me.jpg";

</script>

<main>
	
	<div class="row">
		<div class="leftcolumn">
			<div class="app">👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾👾</div>
		</div>
		<div class="rightcolumn">
			<div class="card">
				<img {src} class="img" height="150px" width="auto" alt="error" />
				<h2>Zelin👾</h2>
				<p>所謂伊人</p>
			</div>
			<div class="card">
				<h3>popular articles</h3>
				<div><p>article</p></div>
				<div><p>article</p></div>
				<div><p>article</p></div>
			</div>
			<div class="card">
				<h3>Follow me</h3>
				<p>author</p>
			</div>
		</div>
		<div class="header">
			<h1>Zelin Design</h1>
			<p>由愛故生憂 由愛故生怖 若離於愛者 無憂亦無怖</p>
		</div>
		<div class="topnav">
			<a>.</a>
		</div>
		<div class="midcolumn">
			<div class="card">
				<h2>文章标题</h2>
				<h5>2019 年 4 月 17日</h5>
				<div class="fakeimg" style="height:200px;">图片</div>
				<p>一些文本...</p>
				<p>菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！</p>
			</div>
			<div class="card">
				<h2>文章标题</h2>
				<h5>2019 年 4 月 17日</h5>
				<div class="fakeimg" style="height:200px;">图片</div>
				<p>一些文本...</p>
				<p>菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！菜鸟教程 - 学的不仅是技术，更是梦想！</p>
			</div>
		</div>
	</div>
</main>

<style>
	* {
	  box-sizing: border-box;
	}
	 
	.app {
		line-height: 180%;
		font-size: 38px;
	}

	/* 头部标题 */
	.header {
	  padding: 8px;
	  text-align: center;
	  color: #9966FF;
	  background: white;
	}
	 
	.header h1 {
	  font-size: 45px;
	}
	 
	/* 导航条 */
	.topnav {
	  overflow: hidden;
	  background-color: rgb(101, 149, 238);
	}
	 
	/* 导航条链接 */
	.topnav a {
	  float: left;
	  display: block;
	  color: #f2f2f2;
	  padding: 1px 16px;
	  text-decoration: none;
	}
	 
	/* 链接颜色修改 */
	.topnav a:hover {
	  background-color: #aaa;
	  color: rgb(160, 26, 154);
	}
	 
	/* 创建两列 */
	/* Left column */
	.leftcolumn {
		float: left;
		height: 100vh;
		padding-left: 10px;
		padding-right: 50px;
		width: 3%;
	}
	.midcolumn {   
	  float: left;
	  width: 72%;
	}
	 
	/* 右侧栏 */
	.rightcolumn {
	  float: right;
	  width: 23%;
	  background-color: #f1f1f1;
	  padding-left: 20px;
	}
	 
	/* 图像部分 */
	.fakeimg {
	  background-color: #aaa;
	  width: 100%;
	  padding: 20px;
	}
	 
	/* 文章卡片效果 */
	.card {
	  background-color: white;
	  padding: 20px;
	  margin-top: 20px;
	}
	 
	/* 列后面清除浮动 */
	.row:after {
	  content: "";
	  display: table;
	  clear: both;
	}
	 
	/* 响应式布局 - 屏幕尺寸小于 800px 时，两列布局改为上下布局 */
	@media screen and (max-width: 800px) {
	  .midcolumn, .rightcolumn {   
		width: 100%;
		padding: 0;
	  }
	}
	 
	/* 响应式布局 -屏幕尺寸小于 400px 时，导航等布局改为上下布局 */
	@media screen and (max-width: 400px) {
	  .topnav a {
		float: none;
		width: 100%;
	  }
	}
</style>

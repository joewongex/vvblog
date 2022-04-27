<script setup></script>

<template>
  <div class="container">
    <div class="header">
      <div class="site-name">
        <a href="/" class="logo">蜗·牛</a>
        <p class="description">My Big Big Park</p>
      </div>
      <div class="nav-menu">
        <a href="/" class="current">
          <span class="svg-icon svg-baseline">
            <font-awesome-icon icon="home" />
          </span>
          <span class="icon-text"> 首页</span>
        </a>
        <a href="/about">
          <span class="svg-icon svg-baseline">
            <font-awesome-icon icon="user" />
          </span>
          <span class="icon-text"> 关于</span>
        </a>
      </div>
    </div>

    <div class="main">
      <div class="left">
        <div class="content-container">
          <div class="post">
            <h1 class="post__title">
              <a href="/emotion/nowhere-to-run.html">无处可逃</a>
            </h1>
            <div class="post__meta">
              <span class="svg-icon svg-baseline">
                <font-awesome-icon icon="calendar-days" />
              </span> <span class="icon-text">2022年03月25日</span>
            </div>
            <div class="post__content">
              已经是被禁足的第14天了，然而解封依然遥遥无期。在3月11日回到家的那个下午，怎么也想不到那竟是最后一次出门。这是国内自初代病毒席卷武汉以来最大规模的一次爆发了，两年前就开始被各种大小疫情袭扰生活的我们，至今依然受此困扰，甚至还变本加厉地身处漩涡的中心，仿佛世界停滞时光倒流，然而无情的时间又怎么可能停止流逝，它已让我的精神和身体以肉眼可见的速度在一天天衰老。
            </div>
            <p class="post__more">
              <a href="/emotion/nowhere-to-run.html">
                阅读全文
                <span class="svg-icon svg-baseline">
                  <font-awesome-icon icon="chevron-right" />
                </span>
              </a>
            </p>
          </div>

          <div class="post">
            <h1 class="post__title">
              <a href="/emotion/nowhere-to-run.html">无处可逃</a>
            </h1>
            <div class="post__meta">
              <span class="svg-icon svg-baseline">
                <font-awesome-icon icon="calendar-days" />
              </span> <span class="icon-text">2022年03月25日</span>
            </div>
            <div class="post__content">
              已经是被禁足的第14天了，然而解封依然遥遥无期。在3月11日回到家的那个下午，怎么也想不到那竟是最后一次出门。这是国内自初代病毒席卷武汉以来最大规模的一次爆发了，两年前就开始被各种大小疫情袭扰生活的我们，至今依然受此困扰，甚至还变本加厉地身处漩涡的中心，仿佛世界停滞时光倒流，然而无情的时间又怎么可能停止流逝，它已让我的精神和身体以肉眼可见的速度在一天天衰老。
            </div>
            <p class="post__more">
              <a href="/emotion/nowhere-to-run.html">
                阅读全文
                <span class="svg-icon svg-baseline">
                  <font-awesome-icon icon="chevron-right" />
                </span>
              </a>
            </p>
          </div>
        </div>
      </div>
      <div class="right">right</div>
      <div class="bottom">bottom</div>
    </div>
  </div>
</template>

<style>
body {
  color: #444;
  font-family: TIBch, "Classic Grotesque W01", "Helvetica Neue", Arial, "Hiragino Sans GB", STHeiti, "Microsoft YaHei", "WenQuanYi Micro Hei", SimSun, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 14px;
}

a {
  text-decoration: none;
}
</style>

<style lang="scss" scoped>
.container {
  padding: 0 60px;
  max-width: 1150px;
  margin: 0 auto;

  .header {
    padding: 58px 0 0;
    text-align: left;
    border-bottom: 1px solid #ddd;
    position: relative;

    .site-name {
      margin-bottom: 40px;

      .logo {
        font: bold 38px/1.12 "Times New Roman", Georgia, Times, sans-serif;
        color: #555;
      }

      .description {
        margin: 0.2em 0 0;
        color: #999;
      }
    }

    .nav-menu {
      margin: 10px 0 -1px;
      padding: 0;
      position: absolute;
      right: 0;
      bottom: 0;

      a {
        display: inline-block;
        padding: 3px 20px 3px;
        line-height: 30px;
        color: #444;
        font-size: 13px;

        &.current {
          border: 1px solid #ddd;
          border-bottom-color: #fff;
        }
      }
    }
  }

  .main {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;

    .left {
      width: 75%;

      .content-container {
        padding-right: 50px;
        padding-top: 20px;

        .post {
          padding: 25px 0 15px;

          &__title {
            margin: 0;
            text-align: left;
            font: bold 25px/1.1 ff-tisa-web-pro, Cambria, "Times New Roman", Georgia, Times, sans-serif;

            a {
              color: #555;
            }
          }

          &__meta {
            padding: 0;
            margin: 15px 0 0;
            color: #6e7173;
            float: left;
            display: inline;
            text-indent: 0.15em;
          }

          &__content {
            clear: left;
            font-size: 15px;
            line-height: 1.77;
            color: #444;
            padding-top: 15px;
            text-align: justify;
            word-break: normal;
          }

          &__more {
            a {
              font-size: 14px;
              color: #444;
              margin: -10px 0;
              padding: 5px 10px;
              border: 1px solid #ddd;
              border-radius: 5px;
              float: right;
            }
          }
        }
      }
    }

    .right {
      width: 25%;
    }
  }
}

.svg-icon {
  display: inline-flex;
  align-self: center;
  position: relative;
  height: 1em;
  width: 1em;
  font-size: 12px;

  svg {
    height: 1em;
    width: 1em;
  }

  &.svg-baseline {
    svg {
      bottom: -0.06em;
      position: absolute;
    }
  }
}
</style>
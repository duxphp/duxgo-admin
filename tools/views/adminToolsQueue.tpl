<template>
  <app-layout>

    <div class="p-4">

      <div class="grid grid-cols-2 gap-4 mb-4">
        <div class="bg-white dark:bg-blackgray-4 rounded shadow">
          <div class="mb-4 text-base p-4">当前队列</div>
          ${ unescape .QueueChart }
        </div>
        <div class="bg-white dark:bg-blackgray-4 rounded shadow">
          <div class="mb-4 text-base p-4">近期队列</div>
          ${ unescape .TaskChart }
        </div>
      </div>

      ${ range.QueueList }
      <div class="bg-white dark:bg-blackgray-4 rounded shadow">

        <div class="flex p-4 items-center text-base">
          队列信息 [${ .Queue }]
        </div>

        <div class="grid grid-cols-4 gap-4 divide-x  divide-gray-200 bg-gray-50">
          <div class="p-6">
            <div class="mb-10 text-gray-600">状态</div>
            <div class="text-lg flex gap-2 items-center">
              ${ if .Paused }
              <icon-pause-circle size="24" class="text-red-600"/>
              已暂停
              ${ else }
              <icon-play-circle size="24" class="text-green-600"/>
              已启动
              ${ end }
            </div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">内存占用 (Redis容量)</div>
            <div class="text-lg">${ .MemoryUsage }</div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">执行中</div>
            <div class="text-lg">${ .Active }</div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">待处理</div>
            <div class="text-lg">${ .Pending }</div>
          </div>
        </div>
        <div class="grid grid-cols-4 gap-4 divide-x  divide-gray-200 bg-gray-50 border-t border-gray-200">
          <div class="p-6">
            <div class="mb-10 text-gray-600">定时任务</div>
            <div class="text-lg">${ .Scheduled }</div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">已完成</div>
            <div class="text-lg">${ .Completed }</div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">处理总数</div>
            <div class="text-lg">${ .ProcessedTotal }</div>
          </div>
          <div class="p-6">
            <div class="mb-10 text-gray-600">失败总数</div>
            <div class="text-lg">${ .FailedTotal }</div>
          </div>
        </div>

      </div>
      ${ end }


    </div>
  </app-layout>
</template>
<script>
export default {
  data() {
    return {

      series: [{
        name: 'PRODUCT A',
        data: [44, 55, 41, 67, 22, 43]
      }, {
        name: 'PRODUCT B',
        data: [13, 23, 20, 8, 13, 27]
      }, {
        name: 'PRODUCT C',
        data: [11, 17, 15, 15, 21, 14]
      }, {
        name: 'PRODUCT D',
        data: [21, 7, 25, 13, 22, 8]
      }],
      chartOptions: {
        chart: {
          type: 'bar',
          height: 350,
          stacked: true,
          toolbar: {
            show: true
          },
          zoom: {
            enabled: true
          }
        },
        responsive: [{
          breakpoint: 480,
          options: {
            legend: {
              position: 'bottom',
              offsetX: -10,
              offsetY: 0
            }
          }
        }],
        plotOptions: {
          bar: {
            horizontal: false,
            borderRadius: 10
          },
        },
        xaxis: {
          type: 'datetime',
          categories: ['01/01/2011 GMT', '01/02/2011 GMT', '01/03/2011 GMT', '01/04/2011 GMT',
            '01/05/2011 GMT', '01/06/2011 GMT'
          ],
        },
        legend: {
          position: 'right',
          offsetY: 40
        },
        fill: {
          opacity: 1
        }
      },


    }
  }
}
</script>
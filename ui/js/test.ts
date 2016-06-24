/// <reference path="../node_modules/zone.js/dist/zone.js.d.ts" />
//function main() {
//  for (let i = 0; i < 10; i++) {
//    recursive(i);
//  }
//}
//
//function recursive(x: number) {
//  for (let i = 0; i < x; i++) {
//    setTimeout(function () {
//      recursive(i-1);
//    }, Math.random() * 5000);
//  }
//}
//
//let zoneSpec: ZoneSpec = {
//  name: 'test',
//  onHasTask(parent: ZoneDelegate, current: Zone, target: Zone, hasTaskState: HasTaskState) {
//    console.log(hasTaskState);
//  },
//  onScheduleTask(parent: ZoneDelegate, current: Zone, target: Zone, task: Task): Task {
//    console.log('schedule', task);
//    current.scheduleTask(
//    return task;
//  },
//  onInvokeTask(parent: ZoneDelegate, current: Zone, target: Zone, task: Task, applyThis: any, applyAny: any) {
//    console.log('invoke', task);
//  }
//}
//
//let testZone = Zone.current.fork(zoneSpec);
//
//
//document.getElementById('start').addEventListener('click', function () {
//  testZone.run(main);
//});

//let zoneSpec: ZoneSpec = {
//  name: 'test',
//  //onScheduleTask(parent: ZoneDelegate, current: Zone, target: Zone, task: Task): Task {
//  //  console.log(task);
//  //}
//  onHasTask(parent: ZoneDelegate, current: Zone, target: Zone, hasTaskState: HasTaskState) {
//    console.log(hasTaskState);
//  }
//}
//let zone = Zone.current.fork(zoneSpec);
//
//let wrapped = zone.wrap(() => {
//  setTimeout(() => {
//    console.log(Zone.current.name);
//  }, 2000);
//}, 'test');
//wrapped();
//
function c(b: number): (c: number) => string {
  return function(c: number): string {
    return "hehe";
  }
}
function Foo(a: (b: number) => (c: number) => string) {
}

() => function(): string {
  return "hehe";
}

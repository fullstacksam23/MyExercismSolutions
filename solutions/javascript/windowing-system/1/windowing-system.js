// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */
export function Size(width=80, height=60) {
  this.width = width;
  this.height = height;
}
Size.prototype.resize = function (newWidth, newHeight) {
  this.width = newWidth;
  this.height = newHeight;
}

export function Position(x=0, y=0){
  this.x = x;
  this.y = y;
}
Position.prototype.move = function(newX, newY){
  this.x = newX;
  this.y = newY;
}

export class ProgramWindow{
  static ScreenHeight = 600;
  static ScreenWidth = 800;
  constructor(){
    this.screenSize = new Size(ProgramWindow.ScreenWidth, ProgramWindow.ScreenHeight);
    this.size = new Size();
    this.position = new Position();
  }
  resize(newSize) {
    const x = this.position.x;
    const y = this.position.y;
    this.size.width = Math.max(1, Math.min(ProgramWindow.ScreenWidth - x, newSize.width));
    this.size.height = Math.max(1,Math.min(ProgramWindow.ScreenHeight - y, newSize.height));
  }
  move(newPosition){
  this.position.x = Math.max(0, Math.min(ProgramWindow.ScreenWidth - this.size.width,newPosition.x));
  this.position.y = Math.max(0, Math.min(ProgramWindow.ScreenHeight - this.size.height,newPosition.y));
  }
}
export function changeWindow(currWindow){
  currWindow.resize(new Size(400, 300));
  currWindow.move(new Position(100, 150));
  return currWindow;
}
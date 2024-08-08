export class BufferWindow {
    constructor(buffer_config){

        this.rows = buffer_config.rows
        this.cols = buffer_config.cols

        this.buffer = [...Array(this.rows)].map(_ => Array(this.cols).fill(' '))
        //this.game_window = $(".game-window")
        this.canvas = document.getElementById("canvas").getContext("2d")
        this.win_height = this.canvas.canvas.height
        this.win_width = this.canvas.canvas.width
    }

    PlacePoint(x, y, char){
        if (x < 0 || x >= this.cols || y < 0 || y >= this.rows) {
            //throw error
        }

        this.buffer[y][x] = char
    }

    RemovePoint(x, y){
        this.buffer[y][x] = ' '
    }

    PlaceVline(x, height, char){
        for (let i = 0; i < height; i++){
            this.buffer[i][x] = char
        }

        for (let i = 0; i < height; i++){
            this.buffer[this.rows-1-i][x] = char
        }
    }

    RemoveVline(x, height){

        console.log("Called RemoveVline")

        for (let i = 0; i < height; i++){
            this.buffer[i][x] = ' '
        }

        for (let i = 0; i < height; i++){
            this.buffer[this.rows-1-i][x] = ' '
        }
    }

    async Render(){

        this.canvas.clearRect(0, 0, this.win_width, this.win_height)

        const pix_height = this.win_height / this.rows
        const pix_width = this.win_width / this.cols

        for (let y = 0; y < this.rows; y ++){
            for (let x = 0; x < this.cols; x ++ ){
                if(this.buffer[y][x] == 'p') {
                    this.canvas.fillStyle = "red"
                }
                else if(this.buffer[y][x] == 'b'){
                    this.canvas.fillStyle = "blue"
                }
                else {
                    continue
                }
                //console.log(`x: ${x} y: ${y} width: ${this.cols} height: ${this.rows}`)
                await this.canvas.fillRect(x*pix_width, y*pix_height, pix_width, pix_height)
            }
        }
    }

    Clear() {
        this.buffer = [...Array(this.rows)].map(_ => Array(this.cols).fill(' '))
    }
}

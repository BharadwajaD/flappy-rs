use ratatui::layout::Rect;

use crate::screen::Screen;

use super::{
    shape::{SRectangle, Shape},
    RenderAble,
};

///(posx, height) is enough to define a pipe
pub struct Pipe {
    pub posx: u16,
    pub nboxes: u16,
    pub boxes: Shape,
}

impl Pipe {
    pub fn new(posx: u16, nboxes: u16) -> Self {

        //TODO: Change these to actual values
        let starty = 0;
        let width = 2; //fixed
        let height = 3; //randomly assigned

        let boxes = Shape::SRect(SRectangle {
            x: posx,
            y: starty ,
            width,
            height: nboxes * height,
        });

        return Self {
            posx,
            nboxes,
            boxes,
        };
    }
}

impl RenderAble for Pipe {
    fn render(&self, screen: &mut dyn Screen) {
        match &self.boxes {
            Shape::SRect(srect) => {
                let area = Rect::new(srect.x, srect.y, srect.width, srect.height);
                log::debug!("Rendering pipe: {:?}" , area);
                screen.draw(area).unwrap();
            }
        };
    }
}

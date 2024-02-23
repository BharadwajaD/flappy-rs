use ratatui::layout::Rect;

use crate::screen::Screen;

use super::RenderAble;

pub struct Bird{
    posx: u16, // to move b/w the pipes
    posy: u16, 
}

impl Bird {
    pub fn new(posx: u16, posy: u16) -> Self {
        return Self { posx, posy  };
    }
}

impl RenderAble for Bird {
    fn render(&self, screen: &mut dyn Screen) {
        let area = Rect::new(self.posx, self.posy, 2, 2);
        screen.draw(area).unwrap();
    }
}

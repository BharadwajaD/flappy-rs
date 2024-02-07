use crate::screen::Screen;

use super::RenderAble;

pub struct Pipe{
}

impl Pipe {
    pub fn new() -> Self {
        return Self {  };
    }
}

impl RenderAble for Pipe {
    fn render(&self, screen: &dyn Screen) {
        screen.get_rectangle();
    }
}

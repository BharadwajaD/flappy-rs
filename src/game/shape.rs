pub struct SRectangle{
    pub x: u16,
    pub y: u16,
    pub width: u16,
    pub height: u16
}

pub enum Shape {
    SRect(SRectangle)
}

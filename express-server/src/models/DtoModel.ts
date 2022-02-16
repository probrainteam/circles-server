import { AbstractDtoModel } from "./AbstractDtoModel";

class DtoModel extends AbstractDtoModel {
    constructor(){
        super();
        this.getConnector().destroy();
    }

}
export { DtoModel }
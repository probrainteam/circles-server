import { Router } from 'express';
import  ExampleController  from '../controllers/ExampleController';

const router = Router();

router.use('/test', ExampleController.exampleControll);

export default router;
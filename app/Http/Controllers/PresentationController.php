<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class PresentationController extends Controller
{
    public function index(){
        return view('welcome');
    }

    public function about(){
        return view('about');
    }


}


<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class company extends Model
{
    protected $table = 'companies';
    use HasFactory;
    public $fillable = [
        'company_name',
        'email',
        'phone',     
    ];    

}


